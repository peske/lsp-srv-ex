package lsp_srv_ex

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// ServerStatus represents the status of the server.
type ServerStatus int

const (
	Created ServerStatus = iota
	Initializing
	Initialized
	Shutdown
)

func (ss ServerStatus) String() string {
	switch ss {
	case Created:
		return "Created"
	case Initializing:
		return "Initializing"
	case Initialized:
		return "Initialized"
	case Shutdown:
		return "Shutdown"
	default:
		return fmt.Sprintf("Unknown status %d", ss)
	}
}

type Helper struct {
	statusLock sync.Mutex
	status     ServerStatus
	logger     *zap.Logger

	Cache *Cache
}

func newHelper(cfg *Config, lgr *zap.Logger) *Helper {
	h := &Helper{}
	if cfg != nil && cfg.Caching {
		h.Cache = &Cache{
			logger: lgr.With(zap.String("object", "Cache")),
		}
	}
	if lgr != nil {
		h.logger = lgr.With(zap.String("object", "Helper"))
	}
	return h
}

func (h *Helper) setStatus(status ServerStatus) (err error) {
	h.statusLock.Lock()
	defer h.statusLock.Unlock()

	success := false

	switch status {
	case Initializing:
		success = h.status == Created
	case Initialized:
		success = h.status == Initializing
	case Shutdown:
		success = h.status == Initialized
	}

	if success {
		h.status = status
	} else {
		err = fmt.Errorf("invalid status transition from '%s' to '%s'", h.status, status)
		h.logger.Error("setStatus", zap.Error(err))
	}

	return err
}

// GetStatus returns the current ServerStatus.
func (h *Helper) GetStatus() ServerStatus {
	h.statusLock.Lock()
	defer h.statusLock.Unlock()
	return h.status
}
