package factory

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/scenario"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type Factory struct {
	testers map[string]*tester.Tester

	scenarios map[string]scenario.Scenario
}

func NewFactory() *Factory {
	return &Factory{
		testers:   map[string]*tester.Tester{},
		scenarios: map[string]scenario.Scenario{},
	}
}

func (fac *Factory) AddScenario(name string, scenario scenario.Scenario) {
	_, ok := fac.scenarios[name]
	if ok {
		panic("scenario already added")
	}
	fac.scenarios[name] = scenario
}

func (fac *Factory) AddTester(name string, tester *tester.Tester) {
	_, ok := fac.testers[name]
	if ok {
		panic("tester already added")
	}
	fac.testers[name] = tester
}

func (fac *Factory) Run(b *testing.B) {
	b.Run("np", func(b *testing.B) {
		for sname, scenario := range fac.scenarios {
			b.Run(sname, func(b *testing.B) {
				for tname, tester := range fac.testers {
					b.Run(tname, func(b *testing.B) {
						scenario.Run(b, tester)
					})
				}
			})
		}
	})
	b.Run("p", func(b *testing.B) {
		for sname, scenario := range fac.scenarios {
			b.Run(sname, func(b *testing.B) {
				for tname, tester := range fac.testers {
					b.Run(tname, func(b *testing.B) {
						b.RunParallel(func(pb *testing.PB) {
							scenario.RunParallel(pb, tester)
						})
					})
				}
			})
		}
	})
}
