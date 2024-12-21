package language

type English struct {
}

func (e English) Speak(eventName string) string {
	switch eventName {
	case "RegisterSuccess":
		return "Register Success!"
	case "HappyNewYear":
		return "Happy New Year!"
	case "CancelClasses":
		return "Cancel Classes!"
	case "ScheduleSuccess":
		return "Schedule Success!"
	}

	return ""
}
