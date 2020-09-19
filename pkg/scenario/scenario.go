package scenario

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type Scenario interface {
	Run(b *testing.B, tester *tester.Tester)

	RunParallel(pb *testing.PB, tester *tester.Tester)
}
