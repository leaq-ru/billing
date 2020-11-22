package robokassa

import "time"

func (w Webhook) pollSubIsValid() (err error) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-w.state.done:
			return
		case <-ticker.C:
			if w.state.sub.IsValid() {
				continue
			}

			err = w.Subscribe()
			if err != nil {
				w.logger.Error().Err(err).Send()
				return
			}
		}
	}
}
