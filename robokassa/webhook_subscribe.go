package robokassa

import "github.com/nats-io/stan.go"

const robokassaWebhookSubjectName = "robokassa_webhook"

func (w Webhook) Subscribe() (err error) {
	w.state.sub, err = w.stanConn.QueueSubscribe(
		robokassaWebhookSubjectName,
		w.serviceName,
		w.cb,
		stan.DurableName(robokassaWebhookSubjectName),
		stan.SetManualAckMode(),
	)
	if err != nil {
		return
	}

	w.state.subscribeCalledOK = true
	return
}
