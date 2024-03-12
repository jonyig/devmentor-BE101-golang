package main

import "fmt"

func main() {
	println("git test")
}

type Notifier interface {
	Notify(name string, lang Language)
}

type Language struct {
	En_Us string
	Zh_Tw string
}

type SignUp struct{}
type BookClass struct{}
type CancelClass struct{}

func (n *SignUp) Notify(name string, lang Language) {
	switch lang {
	case Language{En_Us: lang.En_Us}:
		fmt.Println("SignUp")
	case Language{Zh_Tw: lang.Zh_Tw}:
		fmt.Println("註冊")
	}
}
func (n *BookClass) Notify(name string, lang Language) {
	switch lang {
	case Language{En_Us: lang.En_Us}:
		fmt.Println("BookClass")
	case Language{Zh_Tw: lang.Zh_Tw}:
		fmt.Println("預定課程")
	}
}
func (n *CancelClass) Notify(name string, lang Language) {
	switch lang {
	case Language{En_Us: lang.En_Us}:
		fmt.Println("CancelClass")
	case Language{Zh_Tw: lang.Zh_Tw}:
		fmt.Println("取消課程")
	}
}
