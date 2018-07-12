//+build windows

package main

import (
	"log"
	"os/exec"
)

func checkNetworkStackPlatformSpecific() {

	cmd := exec.Command("ipconfig")
	err := cmd.Run()

	if err != nil {
		log.Printf("failed to execute network check: %v", err)
	}
}
