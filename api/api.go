package api

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
)

type Version struct {
	Version string
}

type API interface {
	ID(context.Context) (peer.ID, error)
	Version(context.Context) (Version, error)
}
