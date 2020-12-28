package rule_testing

// AndError calls each test and returns the first error returned, but invocations stop at the first error.
func AndError(tests ...func() error) error {
	for _, test := range tests {
		err := test()
		if err != nil {
			return err
		}
	}
	return nil
}

// OrError calls each test and return the first error returned. All tests are executed
func OrError(tests ...func() error) (firstErrorIfAny error) {
	for _, test := range tests {
		err := test()
		if err != nil && firstErrorIfAny == nil {
			firstErrorIfAny = err
		}
	}
	return
}
