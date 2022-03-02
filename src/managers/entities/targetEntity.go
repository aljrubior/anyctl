package entities

type TargetEntity interface {
	GetId() string
	GetName() string
	GetType() string
}
