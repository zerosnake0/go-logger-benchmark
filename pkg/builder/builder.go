package builder

import (
	"io"

	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
	"github.com/zerosnake0/go-logger-benchmark/pkg/writer"
)

type TimeConfig struct {
	Format string

	ShowDate     bool
	ShowTime     bool
	ShowMicroSec bool
	UseUTC       bool
}

type LogFormat string

const (
	LogFormatUnknown LogFormat = ""
	LogFormatText    LogFormat = "text"
	LogFormatJson    LogFormat = "json"
)

type LogLevel string

const (
	LogLevelUnknown LogLevel = ""
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarn    LogLevel = "warn"
	LogLevelError   LogLevel = "error"
)

type Config struct {
	Output io.Writer
	Time   TimeConfig
	Format LogFormat
	Color  bool
	Level  LogLevel
}

func (cfg *Config) ApplyDefault() {
	if cfg.Output == nil {
		cfg.Output = writer.DefaultWriter
	}
	if cfg.Level == "" {
		cfg.Level = LogLevelDebug
	}
}

type Builder interface {
	Build(cfg *Config) *tester.Tester
}
