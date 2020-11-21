package balance

import "go.mongodb.org/mongo-driver/bson/primitive"

type Balance struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"u,omitempty"`
	Amount uint32             `bson:"a,omitempty"`
}
