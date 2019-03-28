package tracking

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// TrackingOption ...
type TrackingOption func(client *Tracking)

// Reconfigure ...
func (t *Tracking) Reconfigure(options ...TrackingOption) {
	for _, option := range options {
		option(t)
	}
}

// WithConfiguration ...
func WithConfiguration(config *TrackingConfig) TrackingOption {
	return func(t *Tracking) {
		t.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) TrackingOption {
	return func(t *Tracking) {
		t.logger = logger
		t.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) TrackingOption {
	return func(t *Tracking) {
		t.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) TrackingOption {
	return func(t *Tracking) {
		t.pm = mgr
	}
}
