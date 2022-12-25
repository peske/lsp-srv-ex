package lsp_srv_ex

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
)

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
}

func newHelper(lgr *zap.Logger) *Helper {
	return &Helper{
		logger: lgr,
	}
}

func (h *Helper) SetStatus(status ServerStatus) (err error) {
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

func (h *Helper) GetStatus() ServerStatus {
	h.statusLock.Lock()
	defer h.statusLock.Unlock()
	return h.status
}
