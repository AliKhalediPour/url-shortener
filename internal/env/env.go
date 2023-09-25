package env

import "bale-url-shortener/pkg/env"

type Config struct {
	PostgresDb       string `env:"POSTGRES_DB" envDefault:"bale"`
	PostgresUsername string `env:"POSTGRES_USERNAME" envDefault:"bale"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"bale"`
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort     string `env:"POSTGRES_PORT" envDefault:"5432"`

	RedisHost         string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort         string `env:"REDIS_PORT" envDefault:"6379"`
	RedisExpireMinute int    `env:"REDIS_EXPIRE_MINUTE" envDefault:"60"` // in seconds

	Port     int    `env:"PORT" envDefault:"8000"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func Parse() *Config {
	return env.Parse[Config]()
}
