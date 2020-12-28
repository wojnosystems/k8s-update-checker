package rule_testing

import "github.com/wojnosystems/k8s-update-checker/pkg/k8s"
import testifySuite "github.com/stretchr/testify/suite"

type Suite struct {
	testifySuite.Suite
	SetUp    func(p k8s.Type) (err error)
	TearDown func(p k8s.Type) (err error)
}
