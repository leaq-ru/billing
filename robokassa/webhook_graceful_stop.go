package robokassa

func (w Webhook) GracefulStop() {
	err := w.stanConn.Close()
	if err != nil {
		w.logger.Error().Err(err).Send()
	}
	close(w.state.done)
}
