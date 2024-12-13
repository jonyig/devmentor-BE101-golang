package notification

import "fmt"

type Line struct {
}

func (line Line) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending Line notification")
}
