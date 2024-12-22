package language

import "be101/constants"

type ZHTW struct {
}

func (z ZHTW) Speak(eventName string) string {
	switch eventName {
	case constants.RegisterSuccess:
		return "註冊成功！"
	case constants.HappyNewYear:
		return "新年快樂！"
	case constants.CancelClasses:
		return "課程取消！"
	case constants.ScheduleSuccess:
		return "註冊成功！"
	}

	return ""
}
