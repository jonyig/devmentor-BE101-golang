package language

import "be101/constants"

type English struct {
}

func (e English) Speak(eventName string) string {
	switch eventName {
	case constants.RegisterSuccess:
		return "Register Success!"
	case constants.HappyNewYear:
		return "Happy New Year!"
	case constants.CancelClasses:
		return "Cancel Classes!"
	case constants.ScheduleSuccess:
		return "Schedule Success!"
	}

	return ""
}
