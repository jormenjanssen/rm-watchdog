//+build windows

package main

import (
	"fmt"
	"log"
)

// SdNotify sends a specified string to the systemd notification socket.
func SdNotify(state string) error {

	msg := fmt.Sprintf("Notify: %v", state)
	log.Println(msg)

	return nil
}
