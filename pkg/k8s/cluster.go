package k8s

import (
	"github.com/coreos/go-semver/semver"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

type cluster struct {
	clientset *kubernetes.Clientset
}

func (c *cluster) ServerVersion() (v semver.Version, err error) {
	ver, err := c.clientset.ServerVersion()
	if err != nil {
		return
	}

	major, err := strconv.ParseInt(ver.Major, 10, 64)
	if err != nil {
		return
	}
	minor, err := strconv.ParseInt(ver.Minor, 10, 64)
	if err != nil {
		return
	}
	v = semver.Version{
		Major: major,
		Minor: minor,
		// Patch is always zero from the server
	}
	return
}
