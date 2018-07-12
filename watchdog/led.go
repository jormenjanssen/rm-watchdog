package main

import (
	"time"
)

// ManagerGpio type
type ManagerGpio uint

const (
	// LedUpsGreen led ups green
	LedUpsGreen ManagerGpio = 133
	// LedUpsRed led ups red
	LedUpsRed ManagerGpio = 126
)

// BlinkLed turn the led on for a short period then turns is back off
func BlinkLed(pin ManagerGpio, duration time.Duration) error {

	_ = LedOn(pin)
	time.Sleep(duration)
	_ = LedOff(pin)

	return nil
}
