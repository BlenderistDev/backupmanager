package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/BlenderistDev/backupmanager/internal/services/replace"
)

const (
	daysKeep          = 14
	replacerTimeSleep = 5 * time.Second
)

func main() {
	source := os.Getenv("SOURCE_DIR")
	storageDir := os.Getenv("SOURCE_DIR")

	replacer := replace.Replacer{
		SourceDir:  source,
		StorageDir: storageDir,
		DaysKeep:   daysKeep,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			err := replacer.ReplaceOld()
			if err != nil {
				log.Println(fmt.Sprintf("error in replacer: %v", err))
			}
			time.Sleep(replacerTimeSleep)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM)

	<-shutdown
}
