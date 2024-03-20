package user

type User interface {
	Student
	Guest
}

type Student interface{}
type Guest interface{}
