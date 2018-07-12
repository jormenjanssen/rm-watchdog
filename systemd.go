package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// InstallService function
func InstallService(enableCleanup bool) error {
	err := CopyExecuteable()
	if err != nil {
		return err
	}

	writer, err := getServiceWriter()
	if err != nil {
		return err
	}
	defer writer.Close()

	writeServiceFile(writer, enableCleanup)

	return nil
}

// CopyExecuteable copies the executable
func CopyExecuteable() error {

	exec, err := os.Executable()

	if err != nil {
		return err
	}

	exec, err = filepath.EvalSymlinks(exec)

	if err != nil {
		return err
	}

	from, err := os.Open(exec)

	if err != nil {
		log.Fatal(err)
	}

	defer from.Close()

	to, err := os.OpenFile(WatchdogExecutablePath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func writeServiceFile(writer io.Writer, enableCleanup bool) {

	bufWriter := bufio.NewWriter(writer)
	bufWriter.WriteString(fmt.Sprintln("[Unit]"))
	bufWriter.WriteString(fmt.Sprintln("Description=Watchdog daemon"))

	bufWriter.WriteString(fmt.Sprintln("[service]"))

	if enableCleanup {
		bufWriter.WriteString(fmt.Sprintln("ExecStart=/usr/bin/watchdog"))
	} else {
		bufWriter.WriteString(fmt.Sprintln("ExecStart=/usr/bin/watchdog -disable-cleanup"))
	}

	bufWriter.WriteString(fmt.Sprintln("WatchdogSec=60s"))
	bufWriter.WriteString(fmt.Sprintln("Restart=on-failure"))
	bufWriter.WriteString(fmt.Sprintln("StartLimitInterval=3min"))
	bufWriter.WriteString(fmt.Sprintln("StartLimitBurst=3"))
	bufWriter.WriteString(fmt.Sprintln("StartLimitBurst=3"))

	bufWriter.Flush()
}
