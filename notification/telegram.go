package notification

import "fmt"

type Telegram struct {
}

func (telegram Telegram) Send() {
	fmt.Println("Sending Telegram notification! ")
}
