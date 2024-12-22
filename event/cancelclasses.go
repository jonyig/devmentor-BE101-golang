package event

import (
	"be101/constants"
	"be101/notification"
	"be101/person"
)

type CancelClasses struct {
	notifies []notification.NotificationInterface
}

func (e *CancelClasses) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}

func (e *CancelClasses) Trigger(p person.PersonInterface) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Telegram{})

	for _, notify := range e.notifies {
		notify.Send(p, p.Speak(e.GetName()))
	}
}

func (e *CancelClasses) GetName() string {
	return constants.CancelClasses
}
