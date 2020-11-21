package robokassa

type Robokassa struct {
	webhookSecret string
	merchantLogin string
	passwordOne   string
	passwordTwo   string
	isTest        string // ""/"1"
}
