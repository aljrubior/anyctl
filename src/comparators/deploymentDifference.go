package comparators

type DeploymentDifference struct {
	Operator   string
	Depth      int
	KeyName    interface{}
	LeftValue  interface{}
	RightValue interface{}
}
