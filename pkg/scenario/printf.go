package scenario

import (
	"testing"

	"github.com/zerosnake0/go-logger-benchmark/pkg/method"
	"github.com/zerosnake0/go-logger-benchmark/pkg/tester"
)

type PrintfScenario struct {
	Func func(method method.Printf)
}

func (s *PrintfScenario) Skip(tester *tester.Tester) bool {
	return tester.Printf == nil
}

func (s *PrintfScenario) Run(b *testing.B, tester *tester.Tester) {
	f := s.Func
	method := tester.Printf
	for i := 0; i < b.N; i++ {
		f(method)
	}
}

func (s *PrintfScenario) RunParallel(pb *testing.PB, tester *tester.Tester) {
	f := s.Func
	method := tester.Printf
	for pb.Next() {
		f(method)
	}
}
