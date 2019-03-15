package tracking

import (
	"net/http"

	"github.com/joaosoft/manager"
)

func (c *Controller) RegisterRoutes(web manager.IWeb) error {
	return web.AddRoutes(
		manager.NewRoute(http.MethodPost, "/api/v1/tracking/event", c.AddEventHandler),
	)
}
