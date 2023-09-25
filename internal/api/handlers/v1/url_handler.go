package router

import (
	"fmt"

	dtos "bale-url-shortener/internal/api/dtos"
	util_error "bale-url-shortener/internal/api/utils/errors"
	util_validator "bale-url-shortener/internal/api/utils/validator"

	cache "bale-url-shortener/internal/cache"
	db "bale-url-shortener/internal/db"

	fiber "github.com/gofiber/fiber/v2"
)

type UrlHandler interface {
	GetLongUrl(ctx *fiber.Ctx) error
	GenerateShortUrl(ctx *fiber.Ctx) error
}

type urlHandler struct {
	dbHandler    db.DbHandler
	cacheHandler cache.CacheHandler
}

func NewURLHandler(dbHandler db.DbHandler, cacheHandler cache.CacheHandler) UrlHandler {
	return &urlHandler{
		dbHandler:    dbHandler,
		cacheHandler: cacheHandler,
	}
}

func (u *urlHandler) GetLongUrl(ctx *fiber.Ctx) error {
	// try to extract the short url parameter from the request uri
	urlParam := ctx.Params("url", "")

	// check whether the urlParam is empty or not
	if urlParam == "" {
		return util_error.NewBadRequest("please send the short url")
	}

	// first check with cache
	long, err := u.cacheHandler.Get(ctx.Context(), urlParam)

	if err == nil && long != "" {
		// redirect to the cached long url
		return ctx.Redirect(long, 301)
	}

	// try to get the long url from the db
	long, err = u.dbHandler.GetLongUrl(urlParam)

	if err != nil {
		return util_error.NewInternalError("cannot find the url")
	}

	if long == "" {
		return util_error.NewNotFoundError("url with the given short not found")
	}

	// add the long url to the cache
	err = u.cacheHandler.Add(ctx.Context(), urlParam, long)
	if err != nil {
		return util_error.NewInternalError("cannot cache the urls")
	}

	// redirect with the 301 status code to the desired long link
	return ctx.Redirect(long, 301)
}

func (u *urlHandler) GenerateShortUrl(ctx *fiber.Ctx) error {
	add_user_dto := new(dtos.AddUrlDto)

	// map the request body to the dto(data transfer object)
	if err := ctx.BodyParser(add_user_dto); err != nil {
		return util_error.NewBadRequest(err.Error())
	}

	// validate the body
	if err := util_validator.Validate[dtos.AddUrlDto](*add_user_dto); err != nil {
		return util_error.NewBadRequest(fmt.Sprintf("%v", err))
	}

	// generate the short url
	shortUrl, err := u.dbHandler.AddShortURL(add_user_dto.Long)

	if err != nil {
		return util_error.NewBadRequest(err.Error())
	}

	// add the short and long urls to the cache
	err = u.cacheHandler.Add(ctx.Context(), shortUrl, add_user_dto.Long)
	if err != nil {
		return util_error.NewInternalError("cannot cache the urls")
	}

	return ctx.Status(200).JSON(map[string]any{
		"message": "success",
		"data": map[string]any{
			"longUrl":  add_user_dto.Long,
			"shortUrl": shortUrl,
		},
	})
}
