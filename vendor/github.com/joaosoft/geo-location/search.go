package geolocation

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

const (
	operationSearch  = "search"
	operationReverse = "reverse"
)

type SearchResponse []*Place

type Place struct {
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	PlaceRank   int      `json:"place_rank"`
	BoundingBox []string `json:"bounding_box"`
	AddressType string   `json:"address_type,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

type Address struct {
	Road          string `json:"road"`
	Neighbourhood string `json:"neighbourhood"`
	Suburb        string `json:"suburb"`
	City          string `json:"city"`
	StateDistrict string `json:"state_district"`
	State         string `json:"state"`
	Postcode      string `json:"postcode"`
	Country       string `json:"country"`
	CountryCode2A string `json:"country_code_2a"`
	CountryCode3A string `json:"county_code_3a"`
}

type openStreetMapSearchResponse []openStreetMapPlaceSearch
type openStreetMapReverseResponse openStreetMapPlaceReverse

type openStreetMapPlaceSearch struct {
	PlaceID     int                   `json:"place_id"`
	Licence     string                `json:"licence"`
	OsmType     string                `json:"osm_type"`
	OsmID       int                   `json:"osm_id"`
	Lat         string                `json:"lat"`
	Lon         string                `json:"lon"`
	DisplayName string                `json:"display_name"`
	PlaceRank   int                   `json:"place_rank"`
	Category    string                `json:"category"`
	Type        string                `json:"type"`
	Importance  float64               `json:"importance"`
	Address     *openStreetMapAddress `json:"address,omitempty"`
	Boundingbox []string              `json:"boundingbox"`
}

type openStreetMapPlaceReverse struct {
	PlaceID     int                   `json:"place_id"`
	Licence     string                `json:"licence"`
	OsmType     string                `json:"osm_type"`
	OsmID       int                   `json:"osm_id"`
	Lat         string                `json:"lat"`
	Lon         string                `json:"lon"`
	Name        string                `json:"name"`
	DisplayName string                `json:"display_name"`
	PlaceRank   int                   `json:"place_rank"`
	Category    string                `json:"category"`
	Type        string                `json:"type"`
	Importance  float64               `json:"importance"`
	AddressType string                `json:"addresstype"`
	Address     *openStreetMapAddress `json:"address,omitempty"`
	Boundingbox []string              `json:"boundingbox"`
}

type openStreetMapAddress struct {
	Road          string `json:"road"`
	Neighbourhood string `json:"neighbourhood"`
	Suburb        string `json:"suburb"`
	City          string `json:"city"`
	County        string `json:"county"`
	StateDistrict string `json:"state_district"`
	State         string `json:"state"`
	Postcode      string `json:"postcode"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
}

type SearchService struct {
	service   *GeoLocation
	query     map[string]interface{}
	body      []byte
	method    web.Method
	operation string
}

func newSearchService(service *GeoLocation) *SearchService {
	searchService := &SearchService{
		service: service,
		method:  web.MethodGet,
		query:   make(map[string]interface{}),
	}

	searchService.format(formatJsonV2)

	return searchService
}

func (e *SearchService) format(format format) *SearchService {
	e.query["format"] = format
	return e
}

func (e *SearchService) fetchAddressDetails(enable bool) *SearchService {
	if enable {
		e.query["addressdetails"] = 1
	}
	return e
}

func (e *SearchService) Latitude(latitude float64) *SearchService {
	e.query["lat"] = latitude
	return e
}

func (e *SearchService) Longitude(longitude float64) *SearchService {
	e.query["lon"] = longitude
	return e
}

func (e *SearchService) Street(street string) *SearchService {
	e.query["street"] = street
	return e
}

func (e *SearchService) City(city string) *SearchService {
	e.query["city"] = city
	return e
}

func (e *SearchService) Country(country string) *SearchService {
	e.query["country"] = country
	return e
}

func (e *SearchService) State(state string) *SearchService {
	e.query["state"] = state
	return e
}

func (e *SearchService) PostalCode(postalCode string) *SearchService {
	e.query["postalcode"] = postalCode
	return e
}

func (e *SearchService) CountryCodes(countryCodes ...string) *SearchService {
	e.query["countrycodes"] = strings.Join(countryCodes, ",")
	return e
}

func (e *SearchService) Query(query string) *SearchService {
	e.query["q"] = query
	return e
}

func (e *SearchService) Limit(limit int) *SearchService {
	e.query["limit"] = limit
	return e
}

func (e *SearchService) Body(body []byte) *SearchService {
	e.body = body
	return e
}

func (e *SearchService) Search() (SearchResponse, error) {
	e.operation = operationSearch
	e.fetchAddressDetails(true)

	return e.execute()
}

func (e *SearchService) Reverse() (SearchResponse, error) {
	e.operation = operationReverse

	return e.execute()
}

func (e *SearchService) execute() (SearchResponse, error) {
	var query string

	addSeparator := false

	if len(e.query) > 0 {
		query += "?"
	}

	for name, value := range e.query {
		if addSeparator {
			query += "&"
		}

		query += fmt.Sprintf("%s=%s", name, url.PathEscape(fmt.Sprintf("%+v", value)))
		addSeparator = true
	}

	request, err := e.service.webClient.NewRequest(e.method, fmt.Sprintf("%s/%s%s", e.service.config.Api, e.operation, query), web.ContentTypeApplicationJSON, nil)
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	response, err := request.Send()
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	places := make(SearchResponse, 0)

	switch e.operation {
	case operationSearch:
		apiResponse := openStreetMapSearchResponse{}

		if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
			e.service.logger.Error(err)
			return nil, errors.New(errors.ErrorLevel, 0, err)
		}

		var latitude, longitude float64
		for _, place := range apiResponse {
			latitude, _ = strconv.ParseFloat(place.Lat, 64)
			longitude, _ = strconv.ParseFloat(place.Lon, 64)

			var address *Address
			if place.Address != nil {
				address = &Address{
					Road:          place.Address.Road,
					Neighbourhood: place.Address.Neighbourhood,
					Suburb:        place.Address.Suburb,
					City:          place.Address.City,
					StateDistrict: place.Address.StateDistrict,
					State:         place.Address.State,
					Postcode:      place.Address.Postcode,
					Country:       place.Address.Country,
					CountryCode2A: strings.ToUpper(place.Address.CountryCode),
					CountryCode3A: strings.ToUpper(place.Address.Country),
				}
			}
			places = append(places, &Place{
				Latitude:    latitude,
				Longitude:   longitude,
				Name:        place.DisplayName,
				Category:    place.Category,
				Type:        place.Type,
				PlaceRank:   place.PlaceRank,
				Importance:  place.Importance,
				BoundingBox: place.Boundingbox,
				Address:     address,
			})
		}
	case operationReverse:
		apiResponse := openStreetMapReverseResponse{}

		if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
			e.service.logger.Error(err)
			return nil, errors.New(errors.ErrorLevel, 0, err)
		}

		latitude, _ := strconv.ParseFloat(apiResponse.Lat, 64)
		longitude, _ := strconv.ParseFloat(apiResponse.Lon, 64)

		var address *Address
		if apiResponse.Address != nil {
			address = &Address{
				Road:          apiResponse.Address.Road,
				Neighbourhood: apiResponse.Address.Neighbourhood,
				Suburb:        apiResponse.Address.Suburb,
				City:          apiResponse.Address.City,
				StateDistrict: apiResponse.Address.StateDistrict,
				State:         apiResponse.Address.State,
				Postcode:      apiResponse.Address.Postcode,
				Country:       apiResponse.Address.Country,
				CountryCode2A: strings.ToUpper(apiResponse.Address.CountryCode),
				CountryCode3A: strings.ToUpper(apiResponse.Address.Country),
			}
		}

		places = append(places, &Place{
			Latitude:    latitude,
			Longitude:   longitude,
			Name:        apiResponse.Name,
			Category:    apiResponse.Category,
			Type:        apiResponse.Type,
			PlaceRank:   apiResponse.PlaceRank,
			Importance:  apiResponse.Importance,
			AddressType: apiResponse.AddressType,
			Address:     address,
			BoundingBox: apiResponse.Boundingbox,
		})
	default:
		return nil, nil

	}

	return places, nil
}
