package go_logger_benchmark

import (
	"fmt"
	"testing"
	"time"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	"github.com/zerosnake0/go-logger-benchmark/pkg/builder/logrus"
	"github.com/zerosnake0/go-logger-benchmark/pkg/builder/std"
	"github.com/zerosnake0/go-logger-benchmark/pkg/builder/zerolog"
	"github.com/zerosnake0/go-logger-benchmark/pkg/factory"
	"github.com/zerosnake0/go-logger-benchmark/pkg/method"
	"github.com/zerosnake0/go-logger-benchmark/pkg/scenario"
	"github.com/zerosnake0/go-logger-benchmark/pkg/writer"
)

// Test matrix
// - Builder: How Logger is configured
// - Tester: How Logger is called
// - Scenario: What arguments to be called with
func BenchmarkTest(b *testing.B) {
	fac := factory.NewFactory()
	fac.AddScenario("printf", &scenario.PrintfScenario{
		func(method method.Printf) {
			method("hello %s %d %f", "a", 1234, 345.1)
		},
	})

	w := writer.DefaultWriter

	cfgs := map[string]*builder.Config{
		"default": {
			Output: w,
			Time: builder.TimeConfig{
				Format: time.RFC3339,

				ShowDate: true,
				ShowTime: true,
				UseUTC:   true,
			},
		},
		"json": {
			Output: w,
			Format: builder.LogFormatJson,
			Time: builder.TimeConfig{
				Format: time.RFC3339,

				ShowDate: true,
				ShowTime: true,
				UseUTC:   true,
			},
		},
	}

	builders := map[string]builder.Builder{
		"std":     std.Builder(),
		"logrus":  logrus.Builder(),
		"zerolog": zerolog.Builder(),
	}

	for cname, cfg := range cfgs {
		cfg.ApplyDefault()
		for bname, builder := range builders {
			name := fmt.Sprintf("%s/%s", cname, bname)
			tester := builder.Build(cfg)
			if tester != nil {
				fac.AddTester(name, tester)
			}
		}
	}

	fac.Run(b)
}
