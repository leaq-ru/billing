package robokassa

import "errors"

func (w Webhook) Serve() (err error) {
	if !w.state.subscribeCalledOK {
		err = errors.New("you must call Subscribe before Serve")
		return
	}

	return w.pollSubIsValid()
}
