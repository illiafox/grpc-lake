package rabbitmq

type Exchange struct {
	Name       string `env:"RABBITMQ_EXCHANGE_NAME" env-required:"" `
	Kind       string `env:"RABBITMQ_EXCHANGE_KIND" env-required:""`
	Durable    bool   `env:"RABBITMQ_EXCHANGE_DURABLE" env-required:""`
	AutoDelete bool   `env:"RABBITMQ_EXCHANGE_AUTO_DELETE" env-required:""`
	Internal   bool   `env:"RABBITMQ_EXCHANGE_INTERNAL" env-required:"" `
	NoWait     bool   `env:"RABBITMQ_EXCHANGE_NO_WAIT" env-required:"" `
}
