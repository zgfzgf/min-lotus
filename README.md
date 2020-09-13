
# lotus
`cd cmd/lotus/`  
`go build`  
`./lotus`
# lotus-fx
`cmd/lotus-fx`  
`go build`  
`./lotus-fx`

New ******************************** New  
2020/09/13 13:32:42 [Fx] PROVIDE        main.mystruct <= main.NewMyConstruct()  
2020/09/13 13:32:42 [Fx] PROVIDE        *log.Logger <= reflect.makeFuncStub()  
2020/09/13 13:32:42 [Fx] PROVIDE        http.Handler <= main.NewHandler()  
2020/09/13 13:32:42 [Fx] PROVIDE        *http.ServeMux <= main.NewMux()  
2020/09/13 13:32:42 [Fx] PROVIDE        fx.Lifecycle <= go.uber.org/fx.New.func1()  
2020/09/13 13:32:42 [Fx] PROVIDE        fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()  
2020/09/13 13:32:42 [Fx] PROVIDE        fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm() 
2020/09/13 13:32:42 [Fx] INVOKE         main.invokeAnotherFunc()  
Executing NewLogger.  
invokeAnotherFunc...  
2020/09/13 13:32:42 [Fx] INVOKE         main.invokeUseMyconstruct()  
Executing NewMyConstruct.  
invokeUseMyconstruct..  
2020/09/13 13:32:42 [Fx] INVOKE         main.invokeNothingUse()  
invokeNothingUse...  
2020/09/13 13:32:42 [Fx] INVOKE         main.invokeRegister()    
Executing NewHandler.  
Executing NewMux.  
invokeRegiste..
.  
Start ************************** Start  
2020/09/13 13:32:42 [Fx] START          main.NewLogger()  
logger onstart..  
2020/09/13 13:32:42 [Fx] START          main.invokeAnotherFunc()  
Starting invokeAnotherFunc.  
2020/09/13 13:32:42 [Fx] START          main.NewHandler()  
handler onstart..  
2020/09/13 13:32:42 [Fx] START          main.NewMux()  
Starting HTTP server.  
2020/09/13 13:32:42 [Fx] RUNNING  
Got a request.  
Stop ************************ Stop  
2020/09/13 13:32:42 [Fx] STOP           main.NewMux()  
Stopping HTTP server.  
2020/09/13 13:32:42 [Fx] STOP           main.NewHandler()  
handler onstop..  
2020/09/13 13:32:42 [Fx] STOP           main.invokeAnotherFunc()  
Stopping invokeAnotherFunc..  
2020/09/13 13:32:42 [Fx] STOP           main.NewLogger()  
logger onstop..
