package router

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"

	cache "bale-url-shortener/internal/cache"
	db "bale-url-shortener/internal/db"

	v1 "bale-url-shortener/internal/api/handlers/v1"
	util_error "bale-url-shortener/internal/api/utils/errors"
)

type Router interface {
	Init()
	Run()
}

type router struct {
	app          *fiber.App
	dbHandler    db.DbHandler
	cacheHandler cache.CacheHandler
	port         int
}

func NewRouter(dbHandler db.DbHandler, cacheHandler cache.CacheHandler, port int) Router {
	return &router{
		dbHandler:    dbHandler,
		cacheHandler: cacheHandler,
		port:         port,
	}
}

func (r *router) addUrlHandler(groupRouter fiber.Router) {
	urlHandler := v1.NewURLHandler(r.dbHandler, r.cacheHandler)

	groupRouter.Get("/url/:url", urlHandler.GetLongUrl)        // GET	/v1/url/:url
	groupRouter.Post("/generate", urlHandler.GenerateShortUrl) // POST	/v1/generate
}

func (r *router) Init() {
	// initialize the fiber instance
	r.app = fiber.New(fiber.Config{
		ErrorHandler: util_error.ErrorHandler,
	})

	// add middlewares
	r.app.Use(fiber_logger.New())

	// add v1 handlers to /v1/
	v1Group := r.app.Group("/v1")
	r.addUrlHandler(v1Group)
}

func (r *router) Run() {
	err := r.app.Listen(fmt.Sprintf(":%d", r.port))
	if err != nil {
		panic("cannot listen to the given port " + fmt.Sprint(r.port))
	}
}
