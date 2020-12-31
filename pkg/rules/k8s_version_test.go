package rules

import (
	"context"
	testifySuite "github.com/stretchr/testify/suite"
	"github.com/wojnosystems/k8s-update-checker/pkg/mocks"
	"github.com/wojnosystems/k8s-update-checker/pkg/rule_testing"
	"k8s.io/apimachinery/pkg/version"
	"testing"
)

type K8sVersionSuite struct {
	rule_testing.TestSuite
}

func (suite *K8sVersionSuite) TestVersionMatches() {
	suite.WithContext(func(ctx context.Context) {
		r, err := RequiredVersion.Check(ctx, harness.K8s())
		suite.Require().NoError(err)
		suite.False(r.IsUpgradeBlocked)
	})
}

func (suite *K8sVersionSuite) TestVersionMismatch() {
	expectedVersion := version.Info{
		Major: "0",
		Minor: "1",
	}

	discMock := &mocks.DiscoveryMock{}
	discMock.On("ServerVersion").
		Once().
		Return(&expectedVersion, nil)

	k8sMock := &mocks.K8sMock{}
	k8sMock.On("Discovery").
		Once().
		Return(discMock)

	suite.WithContext(func(ctx context.Context) {
		r, _ := RequiredVersion.Check(ctx, k8sMock)
		suite.True(r.IsUpgradeBlocked)
	})

	k8sMock.AssertExpectations(suite.T())
	discMock.AssertExpectations(suite.T())
}

func TestK8sVersion(t *testing.T) {
	testifySuite.Run(t, new(K8sVersionSuite))
}
