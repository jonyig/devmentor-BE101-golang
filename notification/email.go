package notification

import "fmt"

type Email struct {
}

func (email Email) Send(s string) {
	fmt.Println("Sending email notification! ")
}
