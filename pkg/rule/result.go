package rule

type Result uint8

const (
	ResultInvalid  Result = iota // Not a valid enum value (default "0" value of this type)
	ResultOk                     // No problems detected, no actions to take to fix anything for this rule required
	ResultBlocking               // K8s admin must fix this in order to upgrade
	resultEnd                    // Bound to allow for enum looping, no intrinsic value
)
