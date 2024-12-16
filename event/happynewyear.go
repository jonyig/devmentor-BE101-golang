package event

import (
	"be101/notification"
	"be101/person"
)

type HappyNewYear struct {
	notifies []notification.NotificationInterface
}

func (e *HappyNewYear) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}

func (e *HappyNewYear) Trigger(p *person.User) {
	e.SetNotify(notification.Line{})
	for _, notify := range e.notifies {
		if p.Language == "en-US" {
			notify.Send("HappyNewYear")
		} else if p.Language == "zh-TW" {
			notify.Send("新年快樂")
		}
	}
}
