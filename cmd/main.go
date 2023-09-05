package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BlenderistDev/backupmanager/internal/services/delete"
	"github.com/BlenderistDev/backupmanager/internal/services/replace"
)

func main() {
	const (
		daysKeepReplacer  = 14
		daysKeepDeleter   = 14
		replacerTimeSleep = 5 * time.Second
		deleterTimeSleep  = 5 * time.Second
	)

	sourceDir := os.Getenv("SOURCE_DIR")
	storageDir := os.Getenv("STORAGE_DIR")

	go startReplacer(sourceDir, storageDir, daysKeepReplacer, replacerTimeSleep)
	go startDeleter(storageDir, daysKeepDeleter, deleterTimeSleep)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM)

	<-shutdown
}

func startDeleter(storageDir string, daysKeepDeleter int, deleterTimeSleep time.Duration) {
	deleter := delete.Deleter{
		StorageDir: storageDir,
		DaysKeep:   daysKeepDeleter,
	}

	for {
		err := deleter.DeleteOld()
		if err != nil {
			log.Println(fmt.Sprintf("error in deleter: %v", err))
		}
		time.Sleep(deleterTimeSleep)
	}
}

func startReplacer(sourceDir string, storageDir string, daysKeepReplacer int, replacerTimeSleep time.Duration) {
	replacer := replace.Replacer{
		SourceDir:  sourceDir,
		StorageDir: storageDir,
		DaysKeep:   daysKeepReplacer,
	}

	for {
		err := replacer.ReplaceOld()
		if err != nil {
			log.Println(fmt.Sprintf("error in replacer: %v", err))
		}
		time.Sleep(replacerTimeSleep)
	}
}
