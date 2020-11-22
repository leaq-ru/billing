package stan

import (
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"strings"
)

func NewConn(serviceName, clusterID, natsURL string) (stan.Conn, error) {
	return stan.Connect(
		clusterID,
		strings.Join([]string{
			serviceName,
			uuid.New().String(),
		}, "-"),
		stan.NatsURL(natsURL),
	)
}
