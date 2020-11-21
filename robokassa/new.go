package robokassa

func New(
	webhookSecret,
	merchantLogin,
	passwordOne,
	passwordTwo,
	isTest string,
) Robokassa {
	return Robokassa{
		webhookSecret: webhookSecret,
		merchantLogin: merchantLogin,
		passwordOne:   passwordOne,
		passwordTwo:   passwordTwo,
		isTest:        isTest,
	}
}
