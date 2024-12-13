package notification

import "fmt"

type Sms struct {
}

func (sms Sms) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending SMS notification")
}
