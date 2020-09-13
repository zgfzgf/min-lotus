package client

import (
	"github.com/zgfzgf/min-lotus/api"
	"github.com/zgfzgf/min-lotus/rpclib"
)

func NewRPC(addr string) api.API {
	var res api.Struct
	rpclib.NewClient(addr, "Filecoin", &res.Internal)
	return &res
}
