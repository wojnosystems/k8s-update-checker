package rule_testing

import (
	"github.com/wojnosystems/k8s-update-checker/pkg/k8s"
)

type Harnesser interface {
	K8s() k8s.Type
	Close() error
}
