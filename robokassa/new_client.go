package robokassa

func NewClient(
	merchantLogin,
	passwordOne,
	isTest string,
) Client {
	return Client{
		merchantLogin: merchantLogin,
		passwordOne:   passwordOne,
		isTest:        isTest,
	}
}
