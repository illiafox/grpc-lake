package config

import "time"

// // Database connections // //

type MongoDB struct {
	URI        string        `yaml:"uri"        env:"MONGODB_URI" env-required:"true"`
	Database   string        `yaml:"database"   env:"MONGODB_DATABASE"`
	Collection string        `yaml:"collection" env:"MONGODB_COLLECTION"`
	Timeout    time.Duration `yaml:"timeout"    env:"MONGODB_CONNECT_TIMEOUT" env-default:"5s"`
}

type Redis struct {
	Addrs       []string      `yaml:"addrs"        env:"REDIS_ADDRESSES"    env-required:""`
	Password    string        `yaml:"password"     env:"REDIS_PASSWORD"`
	PoolTimeout time.Duration `yaml:"pool_timeout" env:"REDIS_POOL_TIMEOUT"`
	PoolSize    int           `yaml:"pool_size"    env:"REDIS_POOL_SIZE"`
	IdleSize    int           `yaml:"idle_size"    env:"REDIS_IDLE_SIZE"`
}

type RabbitMQ struct {
	URI string `yaml:"uri" env:"RABBITMQ_URI" env-required:""`

	Key      string `yaml:"key" env:"RABBITMQ_KEY" env-required:""`
	Exchange struct {
		Name       string `yaml:"name"        env:"RABBITMQ_EXCHANGE_NAME"`
		Kind       string `yaml:"kind"        env:"RABBITMQ_EXCHANGE_KIND"`
		Durable    bool   `yaml:"durable"     env:"RABBITMQ_EXCHANGE_DURABLE"`
		AutoDelete bool   `yaml:"auto_delete" env:"RABBITMQ_EXCHANGE_AUTO_DELETE"`
		Internal   bool   `yaml:"internal"    env:"RABBITMQ_EXCHANGE_INTERNAL"`
		NoWait     bool   `yaml:"no_wait"     env:"RABBITMQ_EXCHANGE_NO_WAIT"`
	}
	Queue struct {
		Name       string `yaml:"name"        env:"RABBITMQ_QUEUE_NAME"`
		Durable    bool   `yaml:"durable"     env:"RABBITMQ_QUEUE_DURABLE"`
		AutoDelete bool   `yaml:"auto_delete" env:"RABBITMQ_QUEUE_AUTODELETE"`
		Exclusive  bool   `yaml:"exclusive"   env:"RABBITMQ_QUEUE_EXCLUSIVE"`
		NoWait     bool   `yaml:"no_wait"     env:"RABBITMQ_QUEUE_NOWAIT"`
	}

	PersistentDeliveryMode bool `yaml:"persistent_delivery_mode" env:"RABBITMQ_PERSISTENT_DELIVERY_MODE" env-required:""`
}

// // Local settings // //

type Cache struct {
	CacheExpire time.Duration `yaml:"cache_expire" env:"CACHE_EXPIRE"`
}

type GRPC struct {
	Port int `yaml:"port" env:"GRPC_PORT" env-default:"443"`
}

type HTTP struct {
	Port int `yaml:"port" env:"HTTP_PORT" env-default:"8080"`

	HTTPS struct {
		KeyFile  string `yaml:"key_file"  env:"HTTPS_KEY_FILE_PATH"`
		CertFile string `yaml:"cert_file" env:"HTTPS_CERT_FILE_PATH"`
	}
}

// // Main config // //

type Config struct {
	Flags Flags `yaml:"-"`

	Cache Cache `yaml:"cache"`
	Redis Redis `yaml:"redis"`

	RabbitMQ RabbitMQ `yaml:"rabbitmq"`

	MongoDB MongoDB `yaml:"mongodb"`
	GRPC    GRPC    `yaml:"grpc"`
	HTTP    HTTP    `yaml:"http"`
}
