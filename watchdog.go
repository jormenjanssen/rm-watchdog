package main

import (
	"log"
	"os"
	"time"
)

// UpsLedBlink The gpio for led blink
const UpsLedBlink = 134

// UpsFailLedBlink The gpio for led blik on failure
const UpsFailLedBlink = 135

// WatchdogCheck func
func WatchdogCheck() {

	BlinkLed(UpsLedBlink, 2500*time.Millisecond)
	softLockup()
	CheckNetwork()
}

func softLockup() {

	_, err := os.Stat("/var/run/watchdog.lockup")
	if err == nil {
		log.Println("Soft lockup request performing now be prepared for watchdog system reset ...")

		for {
			BlinkLed(UpsFailLedBlink, 1*time.Second)
		}
	}

}
