package language

type Language interface {
	GetMessage(event_name string) string
}

type ZnTW struct{}
type EnUS struct{}

func (z ZnTW) GetMessage(event_name string) string {
	switch event_name {
	case "registerSuccess":
		return "註冊成功。歡迎加入我們的社群！"
	case "bookingSuccess":
		return "課程預定成功。期待見到你！"
	case "cancellationSuccess":
		return "課程取消成功。希望在其他課程見到你。"

	default:
		return "沒有這個事件"
	}
}

func (e EnUS) GetMessage(event_name string) string {
	switch event_name {
	case "registerSuccess":
		return "Registration successful. Welcome to our community!"
	case "bookingSuccess":
		return "Course successfully booked. We look forward to seeing you!"
	case "cancellationSuccess":
		return "Course successfully cancelled. We hope to see you in other courses."

	default:
		return "No such event"
	}
}

/*
{
    "registerSuccess": {
        "en-us": "Registration successful. Welcome to our community!",
        "zh-tw": "註冊成功。歡迎加入我們的社群！"
    },
    "bookingSuccess": {
        "en-us": "Course successfully booked. We look forward to seeing you!",
        "zh-tw": "課程預定成功。期待見到你！"
     },
    "cancellationSuccess": {
        "en-us": "Course successfully cancelled. We hope to see you in other courses.",
        "zh-tw": "課程取消成功。希望在其他課程見到你。"
     },
}
*/
