package notification

import "fmt"

type Telegram struct {
}

func (telegram Telegram) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending Telegram notification")
}
