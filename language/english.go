package language

type English struct {
}

func (e English) Speak(eventName string) string {
	switch eventName {
	// case event.RegisterSuccess:
	// 	fmt.Printf("註冊成功！")
	case "HappyNewYear":
		return "Happy New Year!"
		// case CancelClasses:
		// 	fmt.Printf("課程取消！")
		// case ScheduleSuccess:
		// 	fmt.Printf("預約課程成功！")
	}

	return ""
}
