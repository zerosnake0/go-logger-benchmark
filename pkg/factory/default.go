package factory

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	"github.com/zerosnake0/go-logger-benchmark/pkg/scenario"
)

var (
	defaultFactory = NewFactory()
)

func AddBuilder(name string, builder builder.Builder) {
	defaultFactory.AddBuilder(name, builder)
}

func AddConfig(name string, cfg *builder.Config) {
	defaultFactory.AddConfig(name, cfg)
}

func AddScenario(name string, scenario scenario.Scenario) {
	defaultFactory.AddScenario(name, scenario)
}

func Run(b *testing.B) {
	defaultFactory.Run(b)
}
