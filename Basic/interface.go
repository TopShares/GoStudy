package Basic

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() string {
	fmt.Println("Connect:", pc.name)
}

func Disconnct(usb interface{}) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconncted:", pc.name)
		return
	} else {
		fmt.Println("Unknown decive.")
	}
	//switch
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnct:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnct(a)
}
