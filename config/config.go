package config

type Config struct {
	ServiceName string
	Grpc        grpc
	MongoDB     mongodb
	STAN        stan
	NATS        nats
	Robokassa   robokassa
	Service     service
	LogLevel    string `envconfig:"LOGLEVEL"`
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type mongodb struct {
	URL string `envconfig:"MONGODB_URL"`
}

type service struct {
	Parser string `envconfig:"SERVICE_PARSER"`
}

type stan struct {
	ClusterID string `envconfig:"STAN_CLUSTERID"`
}

type nats struct {
	URL string `envconfig:"NATS_URL"`
}

type robokassa struct {
	WebhookSecret string `envconfig:"ROBOKASSA_WEBHOOKSECRET"`
	MerchantLogin string `envconfig:"ROBOKASSA_MERCHANTLOGIN"`
	PasswordOne   string `envconfig:"ROBOKASSA_PASSWORDONE"`
	PasswordTwo   string `envconfig:"ROBOKASSA_PASSWORDTWO"`
	IsTest        string `envconfig:"ROBOKASSA_ISTEST"`
}
