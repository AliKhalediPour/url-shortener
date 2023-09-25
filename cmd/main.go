package main

import (
	api "bale-url-shortener/internal/api"
	cache "bale-url-shortener/internal/cache"
	db "bale-url-shortener/internal/db"
	env "bale-url-shortener/internal/env"
	"bale-url-shortener/pkg/logger"
	"fmt"

	"github.com/rs/zerolog"
)

var (
	cfg          *env.Config
	r            api.Router
	log          *zerolog.Logger
	dbHandler    db.DbHandler
	cacheHandler cache.CacheHandler
)

func init() {
	cfg = env.Parse()

	fmt.Printf("cfg: %v\n", *cfg)

	log = logger.NewLogger(cfg.LogLevel)

	// initialize the db handler
	dbHandler = db.NewDbHandler(cfg.PostgresHost, cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDb, cfg.PostgresPort, log)

	// initialize the cache handler
	cacheHandler = cache.NewCacheHandler(cfg.RedisHost, cfg.RedisPort, cfg.RedisExpireMinute, log)

	r = api.NewRouter(dbHandler, cacheHandler, cfg.Port)
}

func main() {
	r.Init()
	r.Run()
}
