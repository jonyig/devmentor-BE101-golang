package language

import (
	"be101/event"
	"fmt"
)

type languagezhTW struct {
}

func (l languagezhTW) Speak(event event.EventInterface) {
	switch event {
	case RegisterSuccess:
		fmt.Printf("註冊成功！")
	case HappyNewYear:
		fmt.Printf("新年快樂！")
	case CancelClasses:
		fmt.Printf("課程取消！")
	case ScheduleSuccess:
		fmt.Printf("預約課程成功！")
	}
}
