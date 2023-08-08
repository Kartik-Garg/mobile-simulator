package main

import (
	"fmt"
	"time"
)

type MobilePhone struct {
	isOn    bool
	battery int
}

func NewMobilePhone() *MobilePhone {
	return &MobilePhone{
		isOn:    false,
		battery: 100,
	}
}

func (mp *MobilePhone) PowerOn() {
	if !mp.isOn {
		mp.isOn = true
		fmt.Println("Phone powered on")
	}
}

func (mp *MobilePhone) PowerOff() {
	if mp.isOn {
		mp.isOn = false
		fmt.Println("Phone powered off")
	}
}

func (mp *MobilePhone) MakeCall(number string) {
	if mp.isOn && mp.battery > 0 {
		fmt.Printf("Calling %s...\n", number)
		mp.battery -= 5
	} else {
		fmt.Println("Cannot make a call. Phone is off or battery is low.")
	}
}

func main() {
	phone := NewMobilePhone()

	phone.PowerOn()
	time.Sleep(time.Second)
	phone.MakeCall("1234567890")

	time.Sleep(time.Second)
	phone.PowerOff()
}
