//+build windows

package main

import (
	"fmt"
	"log"
)

//LedOn function turns led off
func LedOn(pin ManagerGpio) error {
	msg := fmt.Sprintf("Led on: %v", pin)
	log.Println(msg)

	return nil
}

//LedOff function turns led off
func LedOff(pin ManagerGpio) error {
	msg := fmt.Sprintf("Led off: %v", pin)
	log.Println(msg)

	return nil
}
