package robokassa

import "go.mongodb.org/mongo-driver/bson/primitive"

type webhookMessage struct {
	ID             primitive.ObjectID `json:"i"`
	InvID          uint64             `json:"ii"`
	OutSum         float32            `json:"o"`
	SignatureValue string             `json:"s"`
}
