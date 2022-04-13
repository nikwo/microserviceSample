package internal

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"time"
)

var loggerInstance *zerolog.Logger

func initLogger() error {
	level := viper.GetString("log_level")

	logPath := func() string {
		if viper.GetString("log_path") == "" {
			opSystem := runtime.GOOS
			switch opSystem {
			case "linux":
				return "/var/log/gc-forms.log"
			case "windows":
				return "c:\\temp\\gc-forms.log"
			}
			return "gc-forms.log"
		}
		return viper.GetString("log_path")
	}

	file, err := os.OpenFile(logPath(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		err := fmt.Errorf("cannot open the log file %s, Reason : %s", logPath(), err.Error())
		return err
	}

	writer := diode.NewWriter(file, 10000, 10*time.Microsecond, func(missed int) {
		fmt.Printf("Logger dropped %d messages", missed)
	})

	zLogger := zerolog.New(writer).With().Caller().Timestamp().Logger().Output(file)
	loggerInstance = &zLogger

	switch level {
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break
	case "WARNING":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		break
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	return nil
}

func InitLogger() (zerolog.Logger, error) {
	if err := initLogger(); err != nil {
		return zerolog.Logger{}, err
	}
	loggerInstance.Info().Msg("Starting logger...")
	return *loggerInstance, nil
}
