package daemon

import (
	"net/http"

	"github.com/zgfzgf/min-lotus/api"
	"github.com/zgfzgf/min-lotus/rpclib"
)

func serveRPC(api api.API) error {
	rpcServer := rpclib.NewServer()
	rpcServer.Register("Filecoin", api)
	http.Handle("/rpc/v0", rpcServer)
	return http.ListenAndServe(":1234", http.DefaultServeMux)
}
