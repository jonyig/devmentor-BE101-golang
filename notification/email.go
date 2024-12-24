package notification

import (
	"be101/person"
	"fmt"
)

type Email struct {
	MailAddress string
}

func (email Email) Send(p person.PersonInterface, s string) {
	fmt.Println(p.GetPersonName() + " " + s + " by Sending email notification! ")
}
