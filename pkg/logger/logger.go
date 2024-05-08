package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "ts"
	zerolog.LevelFieldName = "lvl"
	zerolog.MessageFieldName = "msg"

	logLevel := GetLogLevel()
	fmt.Println("Logging Level:", logLevel)
	switch logLevel {
	case "Info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		fmt.Println("Info Level:", logLevel)
	case "Warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		fmt.Println("Warn Level:", logLevel)
	case "Error":
		fmt.Println("Error Level:", logLevel)
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "Debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		fmt.Println("Debug Level:", logLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		fmt.Println("default Level:", logLevel)
	}
	Log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true}).With().Timestamp().Caller().Logger() // -> Adds caller details
}

func GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}
