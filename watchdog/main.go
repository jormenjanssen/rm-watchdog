package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

func main() {
	installExecutable := flag.Bool("install-bin", false, fmt.Sprintf("install current executable to: %v", WatchdogExecutablePath))
	installService := flag.Bool("install-service", false, "install executing assembly as systemd service unit (watchdog.service)")
	patchSystemd := flag.Bool("enable-system-watchdog", false, "patch the systemd system file (/etc/systemd/system.conf)")
	disableCleanup := flag.Bool("disable-cleanup", false, fmt.Sprintf("disables removing trace file on cleanup at: %v", TraceDir))
	WatchdogExecutablePath = *flag.String("bin-path", path.Clean(WatchdogExecutablePath), "specifies executable path for installation")

	flag.Parse()

	if !*disableCleanup {
		RemoveOldSignOfLifeTraces(TraceDir)
	}

	if *installService == true {
		log.Println("Installing watchdog as Systemd service unit")
		err := InstallService(!*disableCleanup)
		if err != nil {
			log.Fatalf("Failed to install: %v", err)
		}
		stop("service install completed")
	}

	if *installExecutable == true {
		log.Println(fmt.Sprintf("Installing watchdog executable in: %v", WatchdogExecutablePath))
		err := CopyExecuteable()
		if err != nil {
			log.Fatalf("failed to copy executeable: %v", err)
		}
		stop("stand-alone install completed")
	}

	if *patchSystemd == true {
		log.Println("Patching systemd config to enable watchdog")
	}

	log.Println("Starting watchdog application")
	var readyChan = make(chan int)

	// Main watchdog loop
	go func(c chan int) {
		initial := true
		tick := time.Tick(20 * time.Second)

		for {
			<-tick
			WatchdogCheck()
			Trace()
			Watchdog()

			if initial {
				initial = false
				readyChan <- 0
			}
		}
	}(readyChan)

	// Signal only when we reported ready
	<-readyChan
	Ready()

	log.Println("Watchdog application is started")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for sig := range sigCh {
		if sig == syscall.SIGHUP {
			reload()
		} else {
			stop(fmt.Sprintf("requested by signal: %v", sig))
			break
		}
	}

	Stopping()
}

func stop(reason string) {
	log.Println(fmt.Sprintf("Stopping reason: %v", reason))
	Stopping()

	os.Exit(0)
}

func reload() {

	Reloading()
	log.Println("Nothing to reload")
	Ready()
}
