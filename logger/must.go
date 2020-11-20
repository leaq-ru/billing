package logger

func (l Logger) Must(err error) {
	if err != nil {
		l.ZL.Panic().Err(err).Send()
	}
}
