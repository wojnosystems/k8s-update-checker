package rules

import (
	"context"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/pkg/errors"
	"github.com/wojnosystems/k8s-update-checker/pkg/k8s"
	"github.com/wojnosystems/k8s-update-checker/pkg/rule"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
)

var RequiredVersion = rule.Definition{
	Name:                 "Required Server Version",
	DeprecationReference: "N/A",
	Check: func(ctx context.Context, k kubernetes.Interface) (result rule.Result, err error) {
		var serverInfo *version.Info
		serverInfo, err = k.Discovery().ServerVersion()
		if err != nil {
			return
		}
		version, err := semver.NewVersion(fmt.Sprintf("%s.%s.0", serverInfo.Major, serverInfo.Minor))
		if err != nil {
			err = errors.Wrapf(err, "unable to parse Kubernetes version from server, got: %s", serverInfo.String())
			return
		}

		result = rule.Result{
			IsUpgradeBlocked: !version.Equal(k8s.CurrentVersion),
			HowToFix:         fmt.Sprintf("This version of k8s-update-checker only works if your server's version is: %s", k8s.CurrentVersion.String()),
		}
		return
	},
}
