package rules

import (
	"github.com/stretchr/testify/suite"
	"github.com/wojnosystems/k8s-update-checker/pkg/rule_testing"
	"testing"
)

type Case struct {
	rule_testing.Suite
}

func (suite *Case) TestFindsDprecatedApps() {
	suite.Fail("not implemented")
}

func TestDeprecatedApps(t *testing.T) {
	suite.Run(t, new(Case))
}
