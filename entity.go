package common

type Entity interface {
	GetID() string
	SetID(string)
	GetName() string
	SetName(string)
}
