package billingimpl

import (
	"errors"
	"net/http"
)

var (
	badRequest          = errors.New(http.StatusText(http.StatusBadRequest))
	internalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
)
