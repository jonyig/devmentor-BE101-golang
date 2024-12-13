package main

import (
	"be101/event"
	"be101/person"
)

func main() {
	p := person.User{
		PersonName:   "Alex",
		IdentityType: "User",
		Language:     "zh-TW",
	}
	eventrs := event.RegisterSuccess{}
	eventrs.Trigger(p)
	eventhny := event.HappyNewYear{}
	eventhny.Trigger(p)
}

// func NewEventRegisterSuccess(u *User) EventRegisterSuccess {
// 	// u.IdentityType = "Student"
// 	// newStudent := Student{
// 	// 	Person:    u.Person,
// 	// 	StudentId: fmt.Sprintf("S%06d", rand.Intn(1000000)),
// 	// }
// 	// u.Person = newStudent.Person

// 	return EventRegisterSuccess{
// 		EventID: GenerateUniqueID(),
// 		Person:  newStudent,
// 	}
// }
