package rule_testing

import (
	"k8s.io/client-go/kubernetes"
)

type Harnesser interface {
	K8s() kubernetes.Interface
}

type HarnessStopper interface {
	Harnesser
	Stop() error
}
