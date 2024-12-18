package notification

import "fmt"

type Telegram struct {
}

func (telegram Telegram) Send(s string) {
	fmt.Println("Sending Telegram notification! ")
}
