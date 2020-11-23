package billingimpl

var (
	badRequest          = errors.New(http.StatusText(http.StatusBadRequest))
	internalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
)
