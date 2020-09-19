package std

import (
	"log"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type stdBuilder struct {
}

func Builder() *stdBuilder {
	return &stdBuilder{}
}

func logger(cfg *builder.Config) *log.Logger {
	flag := 0
	if cfg.Time.Format != "" {
		if cfg.Time.ShowDate {
			flag |= log.Ldate
		}
		if cfg.Time.ShowTime {
			flag |= log.Ltime
		}
		if cfg.Time.ShowMicroSec {
			flag |= log.Lmicroseconds
		}
		if cfg.Time.UseUTC {
			flag |= log.LUTC
		}
	}
	logger := log.New(cfg.Output, "", flag)
	return logger
}

func (stdBuilder) Build(cfg *builder.Config) *tester.Tester {
	switch cfg.Format {
	case builder.LogFormatJson:
		return nil
	}
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
		Output: func(depth int, msg string) {
			logger.Output(2, msg)
		},
	}
}
