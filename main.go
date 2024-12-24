package main

import (
	"be101/event"
	"be101/language"
	"be101/person"
)

func main() {
	newuser := person.User{
		PersonName:   "Alex",
		IdentityType: "User",
		Language:     language.ZHTW{},
	}
	newstudent := person.Student{
		PersonName:   "Bob",
		IdentityType: "Student",
		Language:     language.English{},
	}
	event_rs := event.RegisterSuccess{}
	event_rs.Trigger(newstudent)
	event_hny := event.HappyNewYear{}
	event_hny.Trigger(newuser)
	event_cc := event.CancelClasses{}
	event_cc.Trigger(newstudent)
	event_ss := event.ScheduleSuccess{}
	event_ss.Trigger(newuser)
}
