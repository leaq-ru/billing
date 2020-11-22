package robokassa

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (w Webhook) ProcessAsync(
	secret string,
	invID uint64,
	outSum float32,
	signatureValue string,
) (
	err error,
) {
	if secret != w.webhookSecret {
		err = errors.New(http.StatusText(http.StatusBadRequest))
		return
	}

	bytes, err := json.Marshal(webhookMessage{
		ID:             primitive.NewObjectID(),
		InvID:          invID,
		OutSum:         outSum,
		SignatureValue: signatureValue,
	})

	return w.stanConn.Publish(robokassaWebhookSubjectName, bytes)
}
