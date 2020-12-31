package rules

import (
	"github.com/wojnosystems/k8s-update-checker/pkg/k8s"
	"github.com/wojnosystems/k8s-update-checker/pkg/rule_testing"
	"os"
	"testing"
)

var harness rule_testing.Harnesser

func TestMain(m *testing.M) {
	var err error
	harnessStopper, err := k8s.NewMiniKube()
	harness = harnessStopper
	if err != nil {
		panic(err.Error())
	}
	var exitCode int
	func() {
		defer func() {
			_ = harnessStopper.Stop()
		}()
		exitCode = m.Run()
	}()
	os.Exit(exitCode)
}
