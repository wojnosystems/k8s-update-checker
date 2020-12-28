package rules

import (
	"github.com/wojnosystems/k8s-update-checker/pkg/rule_testing"
	"os"
	"testing"
)

var harness rule_testing.Harnesser

func TestMain(m *testing.M) {
	var err error
	harness, err = rule_testing.NewMiniKube("")
	if err != nil {
		panic(err.Error())
	}
	var exitCode int
	func() {
		defer func() {
			_ = harness.Close()
		}()
		exitCode = m.Run()
	}()
	os.Exit(exitCode)
}
