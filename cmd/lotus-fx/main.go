package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

// 这里把main函数放到最前面，可以清晰的先了解代码调用方法，具体函数定义，自行跳跃过去查看
func main() {

	// 整个fx包执行的顺序是
	// 1. 先执行fx.Invoke中的函数列表，按顺序一个一个执行
	// 2. fx.Provide中构造函数，在Invoke需要的时候，再去执行
	// 执行Invoke中的函数时，当前执行的函数传入参数如果用到的变量，则先调用其构造函数

	// 这里构造函数构造出来的变量不需要明显的进行定义，会自动传给invoke函数

	// 比如nothingUserInvoke 这里没有任何传入参数，则在它之前不执行任何构造函数
	//
	// Register执行时，需要mux *http.ServeMux, h http.Handler, logger *log.Logger 三个传入参数，则执行对应的三个构造>函数
	// 但是在执行 NewHandler构造函数时，需要logger，则在其之前执行NewLogger
	//
	// invokeUseMyconstruct 执行时，需要先执行 NewMyConstruct

	// 至于在fx.Lifecycle 中注册的Onstart OnStop 函数，是在app start 之后，按构造函数的顺序来执行，stop时，按相反顺序执行
	fmt.Println("New ******************************** New")
	app := fx.New(

		// 一系列构造函数
		fx.Provide(
			NewMyConstruct,
			as(NewLogger, new(*log.Logger)),
			//NewLogger,
			NewHandler,
			NewMux,
		),

		// 构造函数执行完后，执行初始化函数
		fx.Invoke(invokeAnotherFunc, invokeUseMyconstruct, invokeNothingUse, invokeRegister),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("Start ************************** Start")
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	http.Get("http://localhost:8080/")

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("Stop ************************ Stop")
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

}

// *log.Logger 类型对象的构造函数 （注意：这里指针与非指针类型是严格区分的）
func NewLogger(lc fx.Lifecycle) *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")

	lc.Append(fx.Hook{
		OnStart: func(i context.Context) error {
			logger.Println("logger onstart..")
			return nil
		},
		OnStop: func(i context.Context) error {
			logger.Println("logger onstop..")
			return nil
		},
	})
	return logger
}

// http.Handler 类型对象的构造函数，它输入参数中需要*log.Logger类型，所以在它执行之前，先执行NewLogger
func NewHandler(lc fx.Lifecycle, logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	lc.Append(fx.Hook{
		OnStart: func(i context.Context) error {
			logger.Println("handler onstart..")
			return nil
		},
		OnStop: func(i context.Context) error {
			logger.Println("handler onstop..")
			return nil
		},
	})

	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// *http.ServeMux 类型对象的构造函数，它输入参数中需要*log.Logger类型，所以在它执行之前，先执行NewLogger
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// 自定义类型
type mystruct struct{}

// mystruct 构造函数
func NewMyConstruct(logger *log.Logger) mystruct {
	logger.Println("Executing NewMyConstruct.")
	return mystruct{}
}

// invokeUseMyconstruct 是invoke函数，在其执行之前，需要执行 NewMyConstruct
func invokeUseMyconstruct(logger *log.Logger, c mystruct) {
	logger.Println("invokeUseMyconstruct..")
}
func invokeRegister(h http.Handler, mux *http.ServeMux, logger *log.Logger) {
	logger.Println("invokeRegiste...")
	mux.Handle("/", h)
}

func invokeAnotherFunc(lc fx.Lifecycle, logger *log.Logger) {
	logger.Println("invokeAnotherFunc...")
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting invokeAnotherFunc.")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping invokeAnotherFunc..")
			return nil
		},
	})
}

// 这个invoke函数，不依赖任何输入变量，所以不需要执行任何构造函数
func invokeNothingUse() {
	fmt.Println("invokeNothingUse...")
}

func as(in interface{}, as interface{}) interface{} {
	outType := reflect.TypeOf(as)

	if outType.Kind() != reflect.Ptr {
		panic("outType is not a pointer")
	}
	if reflect.TypeOf(in).Kind() != reflect.Func {
		ctype := reflect.FuncOf(nil, []reflect.Type{outType.Elem()}, false)
		fmt.Println("as 111111111111111111111111 as")
		fmt.Println(ctype)
		fmt.Println(reflect.ValueOf(in))
		return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
			out := reflect.New(outType.Elem())
			out.Elem().Set(reflect.ValueOf(in))

			return []reflect.Value{out.Elem()}
		}).Interface()
	}

	inType := reflect.TypeOf(in)

	ins := make([]reflect.Type, inType.NumIn())
	outs := make([]reflect.Type, inType.NumOut())

	for i := range ins {
		ins[i] = inType.In(i)
	}
	outs[0] = outType.Elem()
	for i := range outs[1:] {
		outs[i+1] = inType.Out(i + 1)
	}

	ctype := reflect.FuncOf(ins, outs, false)
	return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
		outs := reflect.ValueOf(in).Call(args)
		out := reflect.New(outType.Elem())
		out.Elem().Set(outs[0])
		outs[0] = out.Elem()

		return outs
	}).Interface()
}
