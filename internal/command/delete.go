package command

import (
	"github.com/BlenderistDev/backupmanager/internal/service/delete"
)

type del struct {
	args map[string]string
}

// Execute delete cli command
// Deletes old files from directory
// At least one backup will be left for each week
func (d del) Execute() error {
	storageDir, err := parseDir(d.args, storageParam)
	if err != nil {
		return err
	}

	daysKeep, err := parseInt(d.args, storageKeepDaysParam, storageKeepDaysDefault)
	if err != nil {
		return err
	}

	deleter := delete.Deleter{
		StorageDir: storageDir,
		DaysKeep:   daysKeep,
	}

	err = deleter.DeleteOld()
	if err != nil {
		return err
	}

	return nil
}
