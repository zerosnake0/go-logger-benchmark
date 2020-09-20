package tester

import (
	"github.com/zerosnake0/go-logger-benchmark/pkg/method"
)

type Tester struct {
	Print   method.Print
	Println method.Println
	Printf  method.Printf
	Output  method.Output

	Debugf method.Debugf
	Infof  method.Infof
	Warnf  method.Warnf
	Errorf method.Errorf
}
