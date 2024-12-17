package notification

import "fmt"

type Email struct {
}

func (email Email) Send() {
	fmt.Println("Sending email notification! ")
}
