package env

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func Parse[T any]() *T {
	godotenv.Load(".env")

	var cfg T
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("cannot map the environment variables to the config: %s", err.Error()))
	}

	return &cfg
}
