package event_log

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m Model) AlreadyProcessed(ctx context.Context, eventID primitive.ObjectID) (alreadyProcessed bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = m.eventLogs.FindOne(ctx, eventLog{
		ID: eventID,
	}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = nil
			return
		}
		return
	}

	alreadyProcessed = true
	return
}
