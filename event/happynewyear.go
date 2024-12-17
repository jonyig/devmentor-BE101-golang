package event

import (
	"be101/notification"
	"be101/person"
	"fmt"
)

type HappyNewYear struct {
	notifies []notification.NotificationInterface
}

func (e *HappyNewYear) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}

func (e *HappyNewYear) Trigger(p person.User) {
	e.SetNotify(notification.Line{})
	for _, notify := range e.notifies {
		notify.Send()
		if p.Language == "en-US" {
			fmt.Print("Happy New Year! ")
		} else if p.Language == "zh-TW" {
			fmt.Print("新年快樂！")
		}
	}
}
