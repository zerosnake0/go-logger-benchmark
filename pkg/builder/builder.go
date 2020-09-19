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

type Config struct {
	Output io.Writer
	Time   TimeConfig
	Format LogFormat
	Color  bool
}

func (cfg *Config) ApplyDefault() {
	if cfg.Output == nil {
		cfg.Output = writer.DefaultWriter
	}
}

type Builder interface {
	Build(cfg *Config) *tester.Tester
}
