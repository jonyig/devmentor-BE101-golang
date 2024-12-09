package main

import (
	"be101/events"
	"be101/persons"
	"fmt"
)

func main() {
	user1 := persons.User{
		Person: persons.Person{
			PersonName:   "Alex",
			IdentityType: "User",
			Language:     "en-US",
		},
	}

	newEventRegisterSuccess := events.NewEventRegisterSuccess(user1)
	newEventsNewyearCelebrate := events.NewEventNewyearCelebrate(user1)
	user1.Events = append(user1.Events, newEventRegisterSuccess, newEventsNewyearCelebrate)
	eventsList := user1.Events
	for _, event := range eventsList {
		switch e := event.(type) {
		case events.EventRegisterSuccess:
			fmt.Printf("EventRegisterSuccess with ID: %s\n", e.GetEventID())
		case events.EventScheduleSuccess:
			fmt.Printf("EventScheduleSuccess with ID: %s\n", e.GetEventID())
		case events.EventCancelClasses:
			fmt.Printf("EventCancelClasses with ID: %s\n", e.GetEventID())
		case events.EventNewyearCelebrate:
			fmt.Printf("EventNewyearCelebrate with ID: %s\n", e.GetEventID())
		}
	}

}
