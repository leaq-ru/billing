package md

import (
	"context"
	"github.com/nnqq/scr-billing/safeerr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserOID(ctx context.Context) (userID primitive.ObjectID, err error) {
	id, err := GetUserID(ctx)
	if err != nil {
		return
	}

	userID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		err = safeerr.InternalServerError
	}
	return
}
