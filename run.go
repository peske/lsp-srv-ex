package lsp_srv_ex

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
	"github.com/peske/lsp-srv/server"
	"go.uber.org/zap"
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
	if zapLogger == nil {
		if cfg != nil && cfg.ZapConfig != nil {
			logger, err = cfg.ZapConfig.Build()
		} else {
			logger, err = zap.NewProduction()
		}

		if err == nil {
			defer func() {
				_ = logger.Sync()
			}()
		} else {
			return err
		}
	} else {
		logger = zapLogger
	}

	sf := func(clnt protocol.ClientCloser, ctx context.Context, ccl func()) protocol.Server {
		h := newHelper(cfg, logger)
		cw := NewClientWrapper(clnt, h, logger.With(zap.String("object", "clientWrapper")))
		s := serverFactory(cw, ctx, ccl, h)
		return NewServerWrapper(s, h, cfg, logger.With(zap.String("object", "serverWrapper")))
	}

	return server.Run(sf, cfg.toBaseConfig())
}
