package persons

import (
	"be101/events"
	"fmt"
)

type Person struct {
	PersonName   string
	IdentityType string
	Language     string
	Events       []events.Event
}

type User struct {
	Person
}

type Student struct {
	Person
}

func (s Student) Echo() {
	fmt.Println(s.PersonName + " now is a student")
}
