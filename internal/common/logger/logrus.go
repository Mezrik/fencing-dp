package logger

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func Init() {
	SetFormatter(logrus.StandardLogger())

	logrus.SetLevel(logrus.DebugLevel)
}

func SetFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocalEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocalEnv {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
}
