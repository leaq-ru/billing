package md

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func GetUserID(ctx context.Context) (userID string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.New("failed to get metadata")
		return
	}

	val := md.Get("user-id")
	if len(val) != 0 {
		userID = val[0]
	}

	if userID == "" {
		err = errors.New("unauthorized")
	}
	return
}
