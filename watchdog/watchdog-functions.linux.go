//+build linux

package main

import (
	"log"
	"os/exec"
)

func checkNetworkStackPlatformSpecific() {

	cmd := exec.Command("ifconfig")
	err := cmd.Run()

	if err != nil {
		log.Printf("failed to execute network check: %v", err)
	}
}
