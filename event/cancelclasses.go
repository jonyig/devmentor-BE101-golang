package event

import "be101/notification"

type CancelClasses struct {
	notifies []notification.NotificationInterface
}

func (e *CancelClasses) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}
func (e *CancelClasses) Trigger(eventTitle string) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Telegram{})
	for _, notify := range e.notifies {
		notify.Send("CancelClasses")
	}
}
