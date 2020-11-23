package event_log

import "go.mongodb.org/mongo-driver/bson/primitive"

type eventLog struct {
	ID primitive.ObjectID `bson:"_id"`
}
