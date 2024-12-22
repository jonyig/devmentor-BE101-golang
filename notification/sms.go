package notification

import (
	"be101/person"
	"fmt"
)

type Sms struct {
	PhoneNumber string
}

func (sms Sms) Send(p person.PersonInterface, s string) {
	fmt.Println(p.GetPersonName() + " " + s + " by Sending SMS notificatio! ")
}
