package Test

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}
type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connecting:", pc.name)
}

type TVConnector struct {
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("Connected", tv.name)
}

func TestPhone() {
	pc := PhoneConnector{"PhoneConnector"}
	// a = PhoneConnector{"PhoneConnector"}
	var a Connector
	a = Connector(pc)
	// Disconnect(a)
	a.Connect()
	var i interface{}
	fmt.Println(i == nil)
	var p *int = nil
	i = p
	//false
	fmt.Println(i == nil)
}

func Disconnect(usb interface{}) {
	// if pc, ok := usb.(PhoneConnector); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// 	return
	// }
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected", v.name)
	default:
		fmt.Println("Unknown device")
	}
	fmt.Println("Unkown device.")
}
