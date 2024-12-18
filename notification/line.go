package notification

import "fmt"

type Line struct {
}

func (line Line) Send(s string) {
	fmt.Println(s + " by Sending Line notification! ")
}
