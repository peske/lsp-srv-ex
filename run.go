package lsp_srv_ex

import (
	"context"
	"go.uber.org/zap"

	lsp_srv "github.com/peske/lsp-srv"
	"github.com/peske/lsp-srv/lsp/protocol"
)

var logger *zap.Logger

// Run function starts the server.
// Params:
// serverFactory: server factory
// cfg:           Config instance.
// zapLogger:     zap.Logger to use, or nil.
// zapConfig:     Logging configuration to use. It will be ignored if `zapLogger` argument is not nil.
func Run(serverFactory func(protocol.ClientCloser, context.Context, func(), *Helper) protocol.Server, cfg *Config,
	zapLogger *zap.Logger) (err error) {
	if zapLogger != nil {
		logger = zapLogger
	} else if cfg != nil && cfg.ZapConfig != nil {
		logger, err = cfg.ZapConfig.Build()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return err
	}
	defer func() {
		_ = logger.Sync()
	}()

	sf := func(clnt protocol.ClientCloser, ctx context.Context, ccl func()) protocol.Server {
		h := newHelper(logger.Named("Helper"))
		cw := NewClientWrapper(clnt, h, logger.Named("clientWrapper"))
		s := serverFactory(cw, ctx, ccl, h)
		return NewServerWrapper(s, h, cfg, logger.Named("serverWrapper"))
	}

	return lsp_srv.Run(sf, cfg.toBaseConfig())
}
