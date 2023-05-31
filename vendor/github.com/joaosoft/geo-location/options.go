package geolocation

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// GeoLocationOption ...
type GeoLocationOption func(g *GeoLocation)

// Reconfigure ...
func (g *GeoLocation) Reconfigure(options ...GeoLocationOption) {
	for _, option := range options {
		option(g)
	}
}

// WithConfiguration ...
func WithConfiguration(config *GeoLocationConfig) GeoLocationOption {
	return func(g *GeoLocation) {
		g.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) GeoLocationOption {
	return func(g *GeoLocation) {
		g.logger = logger
		g.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) GeoLocationOption {
	return func(g *GeoLocation) {
		g.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) GeoLocationOption {
	return func(g *GeoLocation) {
		g.pm = mgr
	}
}
