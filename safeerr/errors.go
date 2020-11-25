package safeerr

import (
	"errors"
	"net/http"
)

var (
	BadRequest          = errors.New(http.StatusText(http.StatusBadRequest))
	InternalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
	Unauthorized        = errors.New(http.StatusText(http.StatusUnauthorized))
)
