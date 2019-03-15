package tracking

import "github.com/joaosoft/geo-location"

type IStorageDB interface {
	AddEvent(event *Event) error
}

type Interactor struct {
	config      *TrackingConfig
	storage     IStorageDB
	geoLocation *geolocation.GeoLocation
}

func NewInteractor(config *TrackingConfig, storageDB IStorageDB, geoLocation *geolocation.GeoLocation) *Interactor {
	return &Interactor{
		config:      config,
		storage:     storageDB,
		geoLocation: geoLocation,
	}
}

func (i *Interactor) AddEvent(event *Event) (*AddEventResponse, error) {
	log.WithFields(map[string]interface{}{"method": "AddEvent"})
	log.Infof("adding new event [action: %s]", event.Action)

	// load geo-localization
	var searchResponse geolocation.SearchResponse
	var err error

	if event.Latitude != nil && event.Longitude != nil {
		searchResponse, err = i.geoLocation.NewSearch().Latitude(*event.Latitude).Longitude(*event.Longitude).Reverse()
	} else if event.Street != nil {
		searchResponse, err = i.geoLocation.NewSearch().Street(*event.Street).Search()
	}

	if searchResponse != nil && len(searchResponse) > 0 {
		event.Latitude = &searchResponse[0].Latitude
		event.Longitude = &searchResponse[0].Longitude

		if searchResponse[0].Address != nil {
			event.Country = &searchResponse[0].Address.Country
			event.City = &searchResponse[0].Address.City
		}
	}

	if err != nil {
		log.WithFields(map[string]interface{}{"error": err.Error()}).
			Errorf("error getting geo-localization [category: %s, action: %s] %s", event.Category, event.Action, err).ToError()
		return nil, err
	}

	err = i.storage.AddEvent(event)
	if err != nil {
		log.WithFields(map[string]interface{}{"error": err.Error()}).
			Errorf("error adding new event [category: %s, action: %s] %s", event.Category, event.Action, err).ToError()
		return nil, err
	}

	return &AddEventResponse{
		Success: err == nil,
	}, nil
}
