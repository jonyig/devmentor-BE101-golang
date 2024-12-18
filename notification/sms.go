package notification

import "fmt"

type Sms struct {
}

func (sms Sms) Send(s string) {
	fmt.Println("Sending SMS notificatio! ")
}
