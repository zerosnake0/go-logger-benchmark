package factory

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	"github.com/zerosnake0/go-logger-benchmark/pkg/scenario"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type Factory struct {
	// testers map[string]*tester.Tester

	builders map[string]builder.Builder

	configs map[string]*builder.Config

	scenarios map[string]scenario.Scenario
}

func NewFactory() *Factory {
	return &Factory{
		builders:  map[string]builder.Builder{},
		configs:   map[string]*builder.Config{},
		scenarios: map[string]scenario.Scenario{},
	}
}

func (fac *Factory) AddBuilder(name string, builder builder.Builder) {
	_, ok := fac.builders[name]
	if ok {
		panic("builder already added")
	}
	fac.builders[name] = builder
}

func (fac *Factory) AddConfig(name string, cfg *builder.Config) {
	_, ok := fac.configs[name]
	if ok {
		panic("builder config already added")
	}
	fac.configs[name] = cfg
}

func (fac *Factory) AddScenario(name string, scenario scenario.Scenario) {
	_, ok := fac.scenarios[name]
	if ok {
		panic("scenario already added")
	}
	fac.scenarios[name] = scenario
}

type matrixFunc func(b *testing.B, s scenario.Scenario, t *tester.Tester)

func (fac *Factory) matrix(b *testing.B, f matrixFunc) {
	for cname, cfg := range fac.configs {
		cfg.ApplyDefault()
		b.Log(cfg)
		b.Run(cname, func(b *testing.B) {
			for sname, scenario := range fac.scenarios {
				b.Run(sname, func(b *testing.B) {
					for bname, builder := range fac.builders {
						tester := builder.Build(cfg)
						if tester == nil {
							continue
						}
						if scenario.Skip(tester) {
							continue
						}
						b.Run(bname, func(b *testing.B) {
							f(b, scenario, tester)
						})
					}
				})
			}
		})
	}
}

func (fac *Factory) Run(b *testing.B) {
	b.Run("np", func(b *testing.B) {
		fac.matrix(b, func(b *testing.B, s scenario.Scenario, t *tester.Tester) {
			s.Run(b, t)
		})
	})
	b.Run("p", func(b *testing.B) {
		fac.matrix(b, func(b *testing.B, s scenario.Scenario, t *tester.Tester) {
			b.RunParallel(func(pb *testing.PB) {
				s.RunParallel(pb, t)
			})
		})
	})
}
