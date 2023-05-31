package geolocation

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
	"github.com/joaosoft/web"
)

type GeoLocation struct {
	webClient     *web.Client
	config        *GeoLocationConfig
	isLogExternal bool
	logger        logger.ILogger
	pm            *manager.Manager
	mux           sync.Mutex
}

// NewGeoLocation ...
func NewGeoLocation(options ...GeoLocationOption) (*GeoLocation, error) {
	config, simpleConfig, err := newConfig()
	webClient, err := web.NewClient()
	if err != nil {
		return nil, err
	}

	service := &GeoLocation{
		webClient: webClient,
		pm:        manager.NewManager(manager.WithRunInBackground(false)),
		config:    config.GeoLocation,
		logger:    logger.NewLogDefault("geo-location", logger.WarnLevel),
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.GeoLocation != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.GeoLocation.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	service.Reconfigure(options...)

	return service, nil
}

func (e *GeoLocation) NewSearch() *SearchService {
	return newSearchService(e)
}
