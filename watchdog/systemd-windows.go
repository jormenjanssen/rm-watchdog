//+build windows

package main

import (
	"io"
	"os"
)

// WatchdogExecutablePath destination path for deployment
var WatchdogExecutablePath = "C:\\test\\watchdog.exe"

func getServiceWriter() (io.WriteCloser, error) {
	return os.Stdout, nil
}

func installServicePlatformSpecific(unit string) {

}
