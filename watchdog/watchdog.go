package main

import (
	"log"
	"os"
	"time"
)

// WatchdogCheck func
func WatchdogCheck() {
	BlinkLed(LedUpsGreen, 2*time.Second)
	SoftLockup()
	CheckNetwork()
	ValidateNetworkConnectivity()
}

// SoftLockup simulates an software based lockup
func SoftLockup() {
	_, err := os.Stat("/var/run/watchdog.lockup")
	if err == nil {
		log.Println("Soft lockup request performing now be prepared for watchdog system reset ...")

		for {
			BlinkLed(LedUpsGreen, 1*time.Second)
			time.Sleep(1 * time.Second)
		}
	}

}
