package tracking

import (
	"fmt"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/geo-location"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

// AppConfig ...
type AppConfig struct {
	Tracking *TrackingConfig `json:"tracking"`
}

// TrackingConfig ...
type TrackingConfig struct {
	Host              string                         `json:"host"`
	Dbr               *dbr.DbrConfig                 `json:"dbr"`
	TokenKey          string                         `json:"token_key"`
	ExpirationMinutes int64                          `json:"expiration_minutes"`
	Migration         *migration.MigrationConfig     `json:"migration"`
	GeoLocation       *geolocation.GeoLocationConfig `json:"geo-location"`
	Log               struct {
		Level string `json:"level"`
	} `json:"log"`
}

// newConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
