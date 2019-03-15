package tracking

import (
	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// SessionOption ...
type SessionOption func(client *Tracking)

// Reconfigure ...
func (session *Tracking) Reconfigure(options ...SessionOption) {
	for _, option := range options {
		option(session)
	}
}

// WithConfiguration ...
func WithConfiguration(config *TrackingConfig) SessionOption {
	return func(session *Tracking) {
		session.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) SessionOption {
	return func(session *Tracking) {
		log = logger
		session.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) SessionOption {
	return func(session *Tracking) {
		log.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) SessionOption {
	return func(session *Tracking) {
		session.pm = mgr
	}
}
