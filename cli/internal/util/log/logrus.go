package log

import (
	"fmt"
	"os"
	"strings"

	adapter "github.com/bombsimon/logrusr/v2"
	"github.com/go-logr/logr"
	"github.com/sirupsen/logrus"

	"go.datalift.io/datalift/cli/internal/common"
)

const (
	JsonFormat = "json"
	TextFormat = "text"
)

func NewLogrusLogger(fieldLogger logrus.FieldLogger) logr.Logger {
	return adapter.New(fieldLogger, adapter.WithFormatter(func(val interface{}) string {
		return fmt.Sprintf("%v", val)
	}))
}

func NewWithCurrentConfig() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(CreateFormatter(os.Getenv(common.EnvLogFormat)))
	l.SetLevel(createLogLevel())
	return l
}

func CreateFormatter(logFormat string) logrus.Formatter {
	var formatType logrus.Formatter
	switch strings.ToLower(logFormat) {
	case JsonFormat:
		formatType = &logrus.JSONFormatter{}
	case TextFormat:
		if os.Getenv("FORCE_LOG_COLORS") == "1" {
			formatType = &logrus.TextFormatter{ForceColors: true}
		} else {
			formatType = &logrus.TextFormatter{}
		}
	default:
		formatType = &logrus.TextFormatter{}
	}

	return formatType
}

func createLogLevel() logrus.Level {
	level, err := logrus.ParseLevel(os.Getenv(common.EnvLogLevel))
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}
