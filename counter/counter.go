package counter

import "go.mongodb.org/mongo-driver/bson/primitive"

type Counter struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Entity   string             `bson:"e,omitempty"`
	Sequence uint64             `bson:"s,omitempty"`
}
