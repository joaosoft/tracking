package tracking

import (
	"encoding/json"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type Controller struct {
	config     *TrackingConfig
	interactor *Interactor
}

func NewController(config *TrackingConfig, interactor *Interactor) *Controller {
	return &Controller{
		config:     config,
		interactor: interactor,
	}
}

func (c *Controller) AddEventHandler(ctx *web.Context) error {
	request := &AddEventRequest{}

	err := json.Unmarshal(ctx.Request.Body, request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.interactor.AddEvent(&Event{
		IdEvent:   genUI(),
		Category:  request.Category,
		Action:    request.Action,
		Label:     request.Label,
		Value:     request.Value,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Street:    request.Street,
		MetaData:  request.MetaData,
	})
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, response)
}
