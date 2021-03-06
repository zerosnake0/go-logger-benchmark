package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	"github.com/zerosnake0/go-logger-benchmark/pkg/factory"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

func init() {
	factory.AddBuilder("logrus", Builder())
}

type logrusBuilder struct {
}

func Builder() *logrusBuilder {
	return &logrusBuilder{}
}

type formatter struct {
}

var _ logrus.Formatter = &formatter{}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// TODO:
	return nil, nil
}

func logger(cfg *builder.Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(cfg.Output)

	switch cfg.Format {
	case builder.LogFormatJson:
		formatter := &logrus.JSONFormatter{
			DisableHTMLEscape: false,
			PrettyPrint:       false,
		}
		if cfg.Time.Format == "" {
			formatter.DisableTimestamp = true
		} else {
			formatter.DisableTimestamp = false
			formatter.TimestampFormat = cfg.Time.Format
		}
		logger.SetFormatter(formatter)
	default:
		formatter := &logrus.TextFormatter{
			DisableColors:  true,
			DisableQuote:   true,
			DisableSorting: true,
		}
		if cfg.Color {
			formatter.DisableColors = true
		} else {
			formatter.DisableColors = false
		}
		if cfg.Time.Format == "" {
			formatter.DisableTimestamp = true
		} else {
			formatter.DisableTimestamp = false
			formatter.TimestampFormat = cfg.Time.Format
		}
		logger.SetFormatter(formatter)
	}
	switch cfg.Level {
	case builder.LogLevelDebug:
		logger.SetLevel(logrus.DebugLevel)
	case builder.LogLevelInfo:
		logger.SetLevel(logrus.InfoLevel)
	case builder.LogLevelWarn:
		logger.SetLevel(logrus.WarnLevel)
	case builder.LogLevelError:
		logger.SetLevel(logrus.ErrorLevel)
	default:
		panic("no log level")
	}
	return logger
}

func (logrusBuilder) Build(cfg *builder.Config) *tester.Tester {
	logger := logger(cfg)
	return &tester.Tester{
		Print: func(args ...interface{}) {
			logger.Print(args...)
		},
		Println: func(args ...interface{}) {
			logger.Println(args...)
		},
		Printf: func(fmt string, args ...interface{}) {
			logger.Printf(fmt, args...)
		},
		Debugf: func(fmt string, args ...interface{}) {
			logger.Debugf(fmt, args...)
		},
	}
}
