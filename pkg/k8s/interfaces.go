package k8s

import "github.com/coreos/go-semver/semver"

type Type interface {
	ServerVersion() (v semver.Version, err error)
}
