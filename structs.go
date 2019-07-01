package tracking

import (
	"encoding/json"
	"time"

	"github.com/joaosoft/web"
)

type ErrorResponse struct {
	Code    web.Status `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Cause   string     `json:"cause,omitempty"`
}

type AddEventRequest struct {
	Category  *string          `json:"category" validate:"notzero"`
	Action    *string          `json:"action" validate:"notzero"`
	Label     *string          `json:"label"`
	Value     *int64           `json:"value"`
	Viewer    *string          `json:"viewer"`
	Viewed    *string          `json:"viewed"`
	Latitude  *float64         `json:"latitude"`
	Longitude *float64         `json:"longitude"`
	Street    *string          `json:"street"`
	MetaData  *json.RawMessage `json:"meta_data"`
}

type AddEventResponse struct {
	Success bool `json:"success"`
}

type Event struct {
	IdEvent    string    `json:"id_event" db:"id_event"`
	Category   string    `json:"category" db:"-" validate:"notzero"`
	FkCategory string    `json:"-" db:"fk_category"`
	Action     string    `json:"action" db:"-" validate:"notzero"`
	FkAction   string    `json:"-" db:"fk_action"`
	Label      *string   `json:"label" db:"label"`
	Value      *int64    `json:"value" db:"value"`
	Viewer     *string   `json:"viewer" db:"viewer"`
	Viewed     *string   `json:"viewed" db:"viewed"`
	Latitude   *float64  `json:"latitude" db:"latitude"`
	Longitude  *float64  `json:"longitude" db:"longitude"`
	Country    *string   `json:"country" db:"country"`
	City       *string   `json:"city" db:"city"`
	Street     *string   `json:"street" db:"street"`
	MetaData   *string   `json:"meta_data" db:"meta_data"`
	CreatedAt  time.Time `json:"created_at" db.read:"created_at" db.write:"-"`
}

type Category struct {
	IdCategory string    `json:"id_category" db:"id_category"`
	Name       string    `json:"name" db:"name"`
	CreatedAt  time.Time `json:"created_at" db.read:"created_at" db.write:"-"`
}

type Action struct {
	IdAction  string    `json:"id_action" db:"id_action"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db.read:"created_at" db.write:"-"`
}
