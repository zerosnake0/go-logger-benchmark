package go_logger_benchmark

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/zerosnake0/go-logger-benchmark/pkg/builder"
	_ "github.com/zerosnake0/go-logger-benchmark/pkg/builder/logrus"
	_ "github.com/zerosnake0/go-logger-benchmark/pkg/builder/std"
	_ "github.com/zerosnake0/go-logger-benchmark/pkg/builder/zerolog"
	"github.com/zerosnake0/go-logger-benchmark/pkg/factory"
	"github.com/zerosnake0/go-logger-benchmark/pkg/method"
	"github.com/zerosnake0/go-logger-benchmark/pkg/scenario"
	"github.com/zerosnake0/go-logger-benchmark/pkg/writer"
)

var (
	debug bool
)

// Test matrix
// - Builder: How Logger is configured
// - Tester: How Logger is called
// - Scenario: What arguments to be called with
func BenchmarkTest(b *testing.B) {
	for sname, scenario := range map[string]scenario.Scenario{
		"printf": &scenario.PrintfScenario{
			func(method method.Printf) {
				method("hello %s %d %f", "a", 1234, 345.1)
			},
		},
		"debugf": &scenario.DebugfScenario{
			func(method method.Debugf) {
				method("hello %s %d %f", "a", 1234, 345.1)
			},
		},
	} {
		factory.AddScenario(sname, scenario)
	}

	w := writer.DefaultWriter
	if debug {
		w = writer.UniqueWriter(writer.DefaultWriter, func(line []byte) {
			b.Logf("%s", line)
		})
	}
	for name, cfg := range map[string]*builder.Config{
		"default": {
			Output: w,
			Time: builder.TimeConfig{
				Format: time.RFC3339,

				ShowDate: true,
				ShowTime: true,
				UseUTC:   true,
			},
			Level: builder.LogLevelInfo,
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
			Level: builder.LogLevelInfo,
		},
	} {
		factory.AddConfig(name, cfg)
	}
	factory.Run(b)
}

func TestMain(m *testing.M) {
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()
	os.Exit(m.Run())
}
