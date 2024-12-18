package event

import (
	"be101/notification"
	"be101/person"
)

type CancelClasses struct {
	notifies []notification.NotificationInterface
}

func (e *CancelClasses) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}
func (e *CancelClasses) Trigger(p person.User) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Telegram{})
	/*
		for _, notify := range e.notifies {
			notify.Send()
			if p.Language == "en-US" {
				fmt.Printf("Cancel Classes! ")
			} else if p.Language == "zh-TW" {
				fmt.Printf("課程取消！")
			}
		}
	*/
}
