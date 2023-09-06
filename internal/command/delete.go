package command

import (
	"github.com/BlenderistDev/backupmanager/internal/service/delete"
)

type del struct {
	args map[string]string
}

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
