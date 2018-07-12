package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

// Trace function
func Trace() {

	_, err := os.Stat(TraceDir)

	if err != nil {
		os.Mkdir(TraceDir, 0666)
	}

	writeSignOfLife(getFilename(TraceDir))
}

// RemoveOldSignOfLifeTraces function
func RemoveOldSignOfLifeTraces(dir string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Failed to read dir: %v", err)
	}

	for _, f := range files {
		if !f.IsDir() {
			if f.Name() != getFilename(dir) {
				err = os.Remove(path.Join(dir, f.Name()))
				if err != nil {
					log.Printf("Failed to delete old trace: %v", err)
				} else {
					log.Printf("Removed old trace: %v", f.Name())
				}

			}
		}
	}
}

func getFilename(dir string) string {

	date := time.Now()
	trace := fmt.Sprintf("%v%v_%v.trace", date.YearDay(), date.Year(), date.Hour())
	return path.Join(dir, trace)
}

func writeSignOfLife(filename string) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(fmt.Sprintf("Could not write sign of life trace: %v", err))
	}

	defer f.Close()
	aliveMsg := fmt.Sprintln(fmt.Sprintf("SIGN-OF-LIFE: %v", time.Now()))
	_, err = f.WriteString(aliveMsg)
}
