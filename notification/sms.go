package notification

import "fmt"

type Sms struct {
}

func (sms Sms) Send() {
	fmt.Println("Sending SMS notificatio! ")
}
