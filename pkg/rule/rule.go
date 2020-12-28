package rule

import (
	"github.com/wojnosystems/k8s-update-checker/pkg/k8s"
	"net/url"
)

type Type struct {
	Description     string // Describe what this rule evaluates and why it exists
	Logic           func(k k8s.Type)
	Origin          string
	OriginReference url.URL
}
