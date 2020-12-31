package rule_testing

import (
	"context"
	testifySuite "github.com/stretchr/testify/suite"
	"time"
)

type TestSuite struct {
	testifySuite.Suite
}

func (suite *TestSuite) WithContext(callback func(ctx context.Context)) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	callback(ctx)
}
