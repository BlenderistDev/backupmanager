package command

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	deleteservice "github.com/BlenderistDev/backupmanager/internal/service/delete"
	emptyservice "github.com/BlenderistDev/backupmanager/internal/service/emptydir"
	replaceservice "github.com/BlenderistDev/backupmanager/internal/service/replace"
)

type serve struct {
	args map[string]string
}

// Execute serve cli command
// Serves source and storage paths
// with replace, delete and emptydir commands
func (s serve) Execute() error {
	sourceDir, err := parseDir(s.args, sourceParam)
	if err != nil {
		return err
	}

	storageDir, err := parseDir(s.args, storageParam)
	if err != nil {
		return err
	}

	sourceDaysKeep, err := parseInt(s.args, sourceKeepDaysParam, sourceKeepDaysDefault)
	if err != nil {
		return err
	}

	storageDaysKeep, err := parseInt(s.args, storageKeepDaysParam, storageKeepDaysDefault)
	if err != nil {
		return err
	}

	sleep, err := parseInt(s.args, sleepParam, sleepDefault)
	if err != nil {
		return err
	}

	replacer := replaceservice.Replacer{
		SourceDir:  sourceDir,
		StorageDir: storageDir,
		DaysKeep:   sourceDaysKeep,
	}

	deleter := deleteservice.Deleter{
		StorageDir: storageDir,
		DaysKeep:   storageDaysKeep,
	}

	empty := emptyservice.Deleter{StorageDir: storageDir}

	go func() {
		for {
			err := replacer.ReplaceOld()
			if err != nil {
				log.Println("error in replacer", err)
			}
			err = deleter.DeleteOld()
			if err != nil {
				log.Println("error in deleter", err)
			}
			err = empty.DeleteEmptyDirs()
			if err != nil {
				log.Println("error in emptydir", err)
			}

			time.Sleep(time.Duration(sleep) * time.Hour)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM)

	<-shutdown

	return nil
}
