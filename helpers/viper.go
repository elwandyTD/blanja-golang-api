package helpers

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ViperInitConfigLogging() {
	viper.SetConfigFile(".env")
	errConfig := viper.ReadInConfig()
	PanicIfError(errConfig)

	useLoggingFile, err := strconv.ParseBool(ViperGetEnv("LOGGING_FILE"))
	PanicIfError(err)

	if useLoggingFile {
		file, errLog := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0666)
		PanicIfError(errLog)
		logrus.SetOutput(file)
	}
}

func ViperGetEnv(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		panic("Invalid type assertion")
	}

	return value
}
