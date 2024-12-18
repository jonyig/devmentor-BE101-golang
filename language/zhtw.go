package language

type ZHTW struct {
}

func (z ZHTW) Speak(eventName string) string {
	switch eventName {
	// case event.RegisterSuccess:
	// 	fmt.Printf("註冊成功！")
	case "HappyNewYear":
		return "新年快樂！"
		// case CancelClasses:
		// 	fmt.Printf("課程取消！")
		// case ScheduleSuccess:
		// 	fmt.Printf("預約課程成功！")
	}

	return ""
}
