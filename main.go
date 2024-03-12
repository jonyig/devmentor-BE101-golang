package main

import "fmt"

func main() {
	println("git test")
}

type Notifier interface {
	Notify(name string, lang string)
}

type SignUp struct{}
type BookClass struct{}
type CancelClass struct{}

// 沒有要修改就傳實體就好，指針是拿來修改的
func (SignUp) Notify(name string, lang string) {
	switch lang {
	case "en_US":
		fmt.Println(name + " Signed Up")
	case "zh_TW":
		fmt.Println(name + " 已註冊")
	}
}

func (BookClass) Notify(name string, lang string) {
	switch lang {
	case "en_US":
		fmt.Println(name + " Booked a Class")
	case "zh_TW":
		fmt.Println(name + " 已預定課程")
	}
}

func (CancelClass) Notify(name string, lang string) {
	switch lang {
	case "en_US":
		fmt.Println(name + " Canceled a Class")
	case "zh_TW":
		fmt.Println(name + " 已取消課程")
	}
}
