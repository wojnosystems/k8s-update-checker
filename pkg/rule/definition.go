package rule

import (
	"context"
	"k8s.io/client-go/kubernetes"
)

type UpgradeBlockedTester func(ctx context.Context, k kubernetes.Interface) (result Result, err error)

type Definition struct {
	Name                 string
	DeprecationReference string
	Check                UpgradeBlockedTester
}

type Result struct {
	IsUpgradeBlocked bool
	HowToFix         string
}
