package model

type Animal interface {
	GetName() string
	GetColor() string
}

type Human interface {
	GetFirstName() string
	GetLastName() string
}
