//+build linux

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// WatchdogExecutablePath destination path for deployment
var WatchdogExecutablePath = "/usr/bin/watchdog"

func getServiceWriter() (io.WriteCloser, error) {
	return os.OpenFile("/etc/systemd/watchdog.service", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
}

func installServicePlatformSpecific(unit string) error {

	cmdReload := exec.Command("systemctl", "daemon-reload")
	_ = cmdReload.Run()
	if cmdReload.ProcessState.Success() {
		return fmt.Errorf("Failed to reload daemon")
	}

	cmdEnable := exec.Command("systemctl", "enable", unit)
	_ = cmdEnable.Run()
	if cmdEnable.ProcessState.Success() {
		return fmt.Errorf("Failed to install service")
	}

	cmdRestart := exec.Command("systemctl", "restart", unit)
	_ = cmdRestart.Run()

	if cmdRestart.ProcessState.Success() {
		return nil
	}

	return fmt.Errorf("Unkown error while installing service")
}
