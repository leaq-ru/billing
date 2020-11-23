package event_log

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (m Model) Put(ctx context.Context, eventID primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = m.eventLogs.InsertOne(ctx, eventLog{
		ID: eventID,
	})
	return
}
