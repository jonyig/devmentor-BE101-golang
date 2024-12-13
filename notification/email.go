package notification

import "fmt"

type Email struct {
}

func (email Email) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending email notification")
}
