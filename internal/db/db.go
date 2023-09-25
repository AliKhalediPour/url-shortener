package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandler interface {
	IsShortURLExists(shortUrl string) (bool, error)
	GetLongUrl(shortUrl string) (string, error)
	AddShortURL(longUrl string) (string, error)
}

type dbHandler struct {
	g   *gorm.DB
	log *zerolog.Logger
}

func NewDbHandler(host, username, password, dbname, port string, log *zerolog.Logger) DbHandler {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	g, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to postgres db: " + err.Error())
	}

	return &dbHandler{
		g:   g,
		log: log,
	}
}

func (d *dbHandler) IsShortURLExists(shortUrl string) (bool, error) {
	var urlResult Url

	dbRes := d.g.Table(`urls`).
		Where(`short = ?`, shortUrl).
		Find(&urlResult)

	if dbRes.Error != nil {
		d.log.Error().Msgf("error in check the existence of the short url: %s", dbRes.Error)
		return false, dbRes.Error
	}

	if dbRes.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (d *dbHandler) GetLongUrl(shortUrl string) (string, error) {
	var urlResult Url

	// check the table and try to fetch the records by short url
	dbRes := d.g.Table(`urls`).
		Where(`short = ?`, shortUrl).
		Find(&urlResult)

	// check the error returned from the db
	if dbRes.Error != nil {
		d.log.Error().Msgf("error in get long url by short url: %s", dbRes.Error)
		return "", dbRes.Error
	}

	if dbRes.RowsAffected == 0 {
		return "", nil
	}

	return urlResult.Long, nil
}

// AddShortURL tries to add new record to the `url` table by the long url
func (d *dbHandler) AddShortURL(longUrl string) (string, error) {
	// generate random key for the primary key
	id := uuid.NewString()

	// declare the GENERATE_SHORT label and go to this line if the short url generated was added to the db earlier
GENERATE_SHORT:
	// generate the random short url
	short := GenerateRandom()

	// create the url database model and pass the generated fields
	url := Url{
		ID:    fmt.Sprint(id),
		Short: short,
		Long:  longUrl,
	}

	// try to add url to the database
	dbRes := d.g.Table(`urls`).Create(url)

	// check the error returned from the db
	if dbRes.Error != nil {

		pgErr, isOk := dbRes.Error.(*pgconn.PgError)

		// back to the generating random short if it's found on the db to generate new one short url
		if isOk && pgErr.Code == "23505" {
			goto GENERATE_SHORT
		}

		d.log.Error().Msgf("error in add short url: %s", dbRes.Error)
		return "", dbRes.Error
	}

	return url.Short, nil
}
