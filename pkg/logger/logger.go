package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var level map[string]zerolog.Level = map[string]zerolog.Level{
	"debug":    zerolog.DebugLevel,
	"info":     zerolog.InfoLevel,
	"warn":     zerolog.WarnLevel,
	"error":    zerolog.ErrorLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"nolevel":  zerolog.NoLevel,
	"disabled": zerolog.Disabled,
}

func NewLogger(levelStr string) *zerolog.Logger {

	lev, ok := level[levelStr]

	if ok == false {
		lev = zerolog.TraceLevel
	}

	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(lev).
		With().
		Timestamp().
		Caller().
		Logger()

	return &l
}
