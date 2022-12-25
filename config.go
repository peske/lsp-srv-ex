package lsp_srv_ex

import (
	"time"

	lsp_srv "github.com/peske/lsp-srv"
	"go.uber.org/zap"
)

type Config struct {
	// Port on which to run the server.
	Port int

	// Address on which to listen for remote connections. If prefixed by 'unix;', the subsequent address is assumed to
	// be a unix domain socket. Otherwise, TCP is used.
	Address string

	// IdleTimeout - shut down the server when there are no connected clients for this duration.
	IdleTimeout time.Duration

	Caching bool `json:"caching"`

	ZapConfig *zap.Config `json:"zapConfig"`
}

func (c *Config) toBaseConfig() *lsp_srv.Config {
	if c == nil {
		return nil
	}
	return &lsp_srv.Config{
		Port:        c.Port,
		Address:     c.Address,
		IdleTimeout: c.IdleTimeout,
	}
}
