package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	enableCleanup := flag.Bool("enable-cleanup", false, fmt.Sprintf("Disables removing trace file on cleanup at: %v", TraceDir))
	enableTrace := flag.Bool("enable-trace", false, fmt.Sprint("Enables tracing of watchdog data"))

	flag.Parse()

	if *enableCleanup {
		RemoveOldSignOfLifeTraces(TraceDir)
	}

	log.Println("Starting watchdog application")
	var readyChan = make(chan int)

	// Main watchdog loop
	go func(c chan int) {
		initial := true
		tick := time.Tick(20 * time.Second)

		for {
			// Wait for tick.
			<-tick

			// Perform OS watchdog loop
			WatchdogCheck()

			// Trace
			if *enableTrace {
				Trace()
			}

			// Notify the process watchdog
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
			stop(fmt.Sprintf("Requested by signal: %v", sig))
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
