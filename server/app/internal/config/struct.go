package config

import "time"

type MongoDB struct {
	URI        string        `env:"MONGODB_URI"             env-required:""`
	Database   string        `env:"MONGODB_DATABASE"        env-required:""`
	Collection string        `env:"MONGODB_COLLECTION"      env-required:""`
	Timeout    time.Duration `enc:"MONGODB_CONNECT_TIMEOUT" env-default:"5s"`
}

type Redis struct {
	Addrs       []string      `env:"REDIS_ADDRESSES"    env-required:""`
	Password    string        `env:"REDIS_PASSWORD"     env-default:""`
	PoolTimeout time.Duration `env:"REDIS_POOL_TIMEOUT" env-default:"5s"`
	PoolSize    int           `env:"REDIS_POOL_SIZE"    env-default:"-1"`
	IdleSize    int           `env:"REDIS_IDLE_SIZE"    env-default:"-1"`
}

type Cache struct {
	CacheExpire time.Duration `env:"CACHE_EXPIRE" env-required:""`
}

type GRPC struct {
	Port int `env:"GRPC_PORT" env-default:"443"`
}

type HTTP struct {
	Port int `env:"HTTP_PORT" env-default:"8080"`

	HTTPS struct {
		KeyFile  string `env:"HTTPS_KEY_FILE_PATH"`
		CertFile string `env:"HTTPS_CERT_FILE_PATH"`
	}
}

type RabbitMQ struct {
	URI string `env:"RABBITMQ_URI" env-required:""`

	Key      string `env:"RABBITMQ_KEY"      env-required:""`
	Exchange struct {
		Name       string `env:"RABBITMQ_EXCHANGE_NAME" env-required:"" `
		Kind       string `env:"RABBITMQ_EXCHANGE_KIND" env-required:""`
		Durable    bool   `env:"RABBITMQ_EXCHANGE_DURABLE" env-required:""`
		AutoDelete bool   `env:"RABBITMQ_EXCHANGE_AUTO_DELETE" env-required:""`
		Internal   bool   `env:"RABBITMQ_EXCHANGE_INTERNAL" env-required:"" `
		NoWait     bool   `env:"RABBITMQ_EXCHANGE_NO_WAIT" env-required:"" `
	}
	Queue struct {
		Name       string `env:"RABBITMQ_QUEUE_NAME" env-required:""`
		Durable    bool   `env:"RABBITMQ_QUEUE_DURABLE" env-required:""`
		AutoDelete bool   `env:"RABBITMQ_QUEUE_AUTODELETE" env-required:""`
		Exclusive  bool   `env:"RABBITMQ_QUEUE_EXCLUSIVE" env-required:""`
		NoWait     bool   `env:"RABBITMQ_QUEUE_NOWAIT" env-required:""`
	}

	PersistentDeliveryMode bool `env:"RABBITMQ_PERSISTENT_DELIVERY_MODE" env-required:""`
}

type Config struct {
	Flags Flags

	Cache Cache
	Redis Redis

	RabbitMQ RabbitMQ

	MongoDB MongoDB
	GRPC    GRPC
	HTTP    HTTP
}
