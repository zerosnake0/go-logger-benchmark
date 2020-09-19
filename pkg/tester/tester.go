package tester

import (
	"github.com/zerosnake0/go-logger-benchmark/pkg/method"
)

type Tester struct {
	Print   func(args ...interface{})
	Println func(args ...interface{})
	Printf  method.Printf
	Output  func(depth int, msg string)
}
