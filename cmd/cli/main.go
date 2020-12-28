package main

import (
	"flag"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// This is just a test, it's not actually used for anything real
func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	_ = kubeconfig
}
