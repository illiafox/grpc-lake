package config

import "time"

type MongoDB struct {
	URI        string        `env:"MONGODB_URI" env-required:""`
	Database   string        `env:"MONGODB_DATABASE" env-required:""`
	Collection string        `env:"MONGODB_COLLECTION" env-required:""`
	Timeout    time.Duration `enc:"MONGODB_CONNECT_TIMEOUT" env-default:"5s"`
}

type Redis struct {
	Addrs       []string      `env:"REDIS_ADDRESSES"  env-required:""`
	Password    string        `env:"REDIS_PASSWORD" env-default:""`
	PoolTimeout time.Duration `env:"REDIS_POOL_TIMEOUT" env-default:"5s"`
	PoolSize    int           `env:"REDIS_POOL_SIZE" env-default:"-1"`
	IdleSize    int           `env:"REDIS_IDLE_SIZE" env-default:"-1"`
}

type Cache struct {
	CacheExpire time.Duration `env:"CACHE_EXPIRE" env-default:"30m"`
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

type Config struct {
	Cache Cache
	Redis Redis

	MongoDB MongoDB
	GRPC    GRPC
	HTTP    HTTP
}
