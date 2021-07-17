package pagination

import (
	"errors"
	"github.com/leaq-ru/proto/codegen/go/opts"
)

var ErrLimitInvalid = errors.New("limit out of 1-100")

func ApplyDefaultLimit(req opter) (limit uint32, err error) {
	limit = 20
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 {
			err = ErrLimitInvalid
			return
		}

		if req.GetOpts().GetLimit() != 0 {
			limit = req.GetOpts().GetLimit()
		}
	}
	return
}

type opter interface {
	GetOpts() *opts.SkipLimit
}
