package response

type TargetResponse interface {
	GetId() string
	GetName() string
	GetType() string
}
