package lsp_srv_ex

import (
	"context"

	lsp_srv "github.com/peske/lsp-srv"
	"github.com/peske/lsp-srv/lsp/protocol"

	"github.com/peske/lsp-srv-ex/internal"
)

func Run(serverFactory func(protocol.ClientCloser, context.Context, func()) protocol.Server, cfg *Config) error {
	sf := func(clnt protocol.ClientCloser, ctx context.Context, ccl func()) protocol.Server {
		cw := internal.NewClientWrapper(clnt)
		s := serverFactory(cw, ctx, ccl)
		return internal.NewServerWrapper(s)
	}
	return lsp_srv.Run(sf, cfg.toBaseConfig())
}
