package event

import (
	"be101/notification"
	"be101/person"
)

type ScheduleSuccess struct {
	notifies []notification.NotificationInterface
}

func (e *ScheduleSuccess) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}

func (e *ScheduleSuccess) RoleTransition(p person.PersonInterface) person.PersonInterface {
	newStudent := person.Student{
		StudentId: "000",
	}

	return newStudent
}

func (e *ScheduleSuccess) Trigger(eventTitle string) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Telegram{})
	for _, notify := range e.notifies {
		notify.Send("ScheduleSuccess")
	}
}
