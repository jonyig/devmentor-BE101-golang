package notification

import (
	"be101/person"
	"fmt"
)

type Line struct {
}

func (line Line) Send(p person.PersonInterface, s string) {
	fmt.Println(p.GetPersonName() + " " + s + " by Sending Line notification! ")
}
