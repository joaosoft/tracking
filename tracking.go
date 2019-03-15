package tracking

import (
	"sync"

	"github.com/joaosoft/geo-location"
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

type Tracking struct {
	config        *TrackingConfig
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	mux           sync.Mutex
}

// NewTracking ...
func NewTracking(options ...SessionOption) (*Tracking, error) {
	config, simpleConfig, err := NewConfig()

	service := &Tracking{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("tracking", logger.WarnLevel),
		config: config.Tracking,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(logger.Instance))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Tracking != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Tracking.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.Tracking = &TrackingConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	// execute migrations
	migrationService, err := migration.NewCmdService(migration.WithCmdConfiguration(service.config.Migration))
	if err != nil {
		return nil, err
	}

	if _, err := migrationService.Execute(migration.OptionUp, 0, migration.ExecutorModeDatabase); err != nil {
		return nil, err
	}

	web := service.pm.NewSimpleWebServer(config.Tracking.Host)

	storage, err := NewStoragePostgres(config.Tracking)
	if err != nil {
		return nil, err
	}

	geolocation, err := geolocation.NewGeoLocation(geolocation.WithConfiguration(config.Tracking.GeoLocation))
	if err != nil {
		return nil, err
	}

	interactor := NewInteractor(config.Tracking, storage, geolocation)

	controller := NewController(config.Tracking, interactor)
	controller.RegisterRoutes(web)

	service.pm.AddWeb("api_web", web)

	return service, nil
}

// Start ...
func (m *Tracking) Start() error {
	return m.pm.Start()
}

// Stop ...
func (m *Tracking) Stop() error {
	return m.pm.Stop()
}
