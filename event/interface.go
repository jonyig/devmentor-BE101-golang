package event

import "be101/person"

type EventInterface interface {
	Trigger(p person.User) string
	GetName() string
}
