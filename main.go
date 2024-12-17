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
	event_rs := event.RegisterSuccess{}
	event_rs.Trigger(p)
	event_hny := event.HappyNewYear{}
	event_hny.Trigger(p)
	event_cc := event.CancelClasses{}
	event_cc.Trigger(p)
	event_ss := event.ScheduleSuccess{}
	event_ss.Trigger(p)
}
