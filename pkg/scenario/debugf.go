package scenario

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/method"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type DebugfScenario struct {
	Func func(method method.Debugf)
}

func (s *DebugfScenario) Skip(tester *tester.Tester) bool {
	return tester.Debugf == nil
}

func (s *DebugfScenario) Run(b *testing.B, tester *tester.Tester) {
	f := s.Func
	method := tester.Debugf
	for i := 0; i < b.N; i++ {
		f(method)
	}
}

func (s *DebugfScenario) RunParallel(pb *testing.PB, tester *tester.Tester) {
	f := s.Func
	method := tester.Debugf
	for pb.Next() {
		f(method)
	}
}
