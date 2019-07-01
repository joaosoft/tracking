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
	request := AddEventRequest{}

	// parameters on url
	err := ctx.Request.BindParams(&request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	// override parameters on body
	err = ctx.Request.Bind(&request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	var metadata *string
	if request.MetaData != nil {
		metadataBytes, err := json.Marshal(request.MetaData)
		if err != nil {
			return ctx.Response.JSON(web.StatusBadRequest, err)
		}

		if len(metadataBytes) > 0 {
			tmp := string(metadataBytes)
			metadata = &tmp
		}
	}

	response, err := c.interactor.AddEvent(&Event{
		IdEvent:   genUI(),
		Category:  *request.Category,
		Action:    *request.Action,
		Label:     request.Label,
		Value:     request.Value,
		Viewer:    request.Viewer,
		Viewed:    request.Viewed,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Street:    request.Street,
		MetaData:  metadata,
	})
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, response)
}
