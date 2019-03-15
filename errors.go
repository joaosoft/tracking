package tracking

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorNotFound = errors.New(errors.ErrorLevel, int(web.StatusNotFound), "not found")
)
