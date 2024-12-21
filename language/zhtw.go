package language

type ZHTW struct {
}

func (z ZHTW) Speak(eventName string) string {
	switch eventName {
	case "RegisterSuccess":
		return "註冊成功！"
	case "HappyNewYear":
		return "新年快樂！"
	case "CancelClasses":
		return "課程取消！"
	case "ScheduleSuccess":
		return "註冊成功！"
	}

	return ""
}
