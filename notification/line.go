package notification

import "fmt"

type Line struct {
}

func (line Line) Send() {
	fmt.Println("Sending Line notification! ")
}
