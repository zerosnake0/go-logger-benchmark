package zerolog

import (
	"github.com/rs/zerolog"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type zerologBuilder struct {
}

func Builder() *zerologBuilder {
	return &zerologBuilder{}
}

func logger(cfg *builder.Config) *zerolog.Logger {
	var logger zerolog.Logger
	switch cfg.Format {
	case builder.LogFormatJson:
		logger = zerolog.New(cfg.Output)
	default:
		logger = zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = cfg.Output
			if cfg.Time.Format != "" {
				w.TimeFormat = cfg.Time.Format
			}
			if cfg.Color {
				w.NoColor = false
			} else {
				w.NoColor = true
			}
		}))
	}
	if cfg.Time.Format != "" {
		logger = logger.With().Timestamp().Logger()
	}
	return &logger
}

func (b *zerologBuilder) Build(cfg *builder.Config) *tester.Tester {
	logger := logger(cfg)
	return &tester.Tester{
		Printf: func(fmt string, args ...interface{}) {
			logger.Info().Msgf(fmt, args...)
		},
	}
}
