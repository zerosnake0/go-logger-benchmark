package scenario

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type Scenario interface {
	Skip(tester *tester.Tester) bool

	Run(b *testing.B, tester *tester.Tester)

	RunParallel(pb *testing.PB, tester *tester.Tester)
}
