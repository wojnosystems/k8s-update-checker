package k8s

import "github.com/coreos/go-semver/semver"

var (
	// CurrentVersion is the version of K8s that we're running and want to upgrade to the next version
	// Patch is always zero (0) as K8s does not return a patch version.
	CurrentVersion = *semver.Must(semver.NewVersion("1.15.0"))
	ToVersion      semver.Version
)

func init() {
	ToVersion = CurrentVersion
	// Assume we're upgrading to the next version, but may change if K8s does a major upgrade.
	ToVersion.Minor += 1
}
