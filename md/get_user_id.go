package md

import (
	"context"
	"github.com/leaq-ru/billing/safeerr"
	"google.golang.org/grpc/metadata"
)

func GetUserID(ctx context.Context) (userID string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = safeerr.InternalServerError
		return
	}

	val := md.Get("user-id")
	if len(val) != 0 {
		userID = val[0]
	}

	if userID == "" {
		err = safeerr.Unauthorized
	}
	return
}
