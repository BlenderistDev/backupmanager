package command

import (
	service "github.com/BlenderistDev/backupmanager/internal/service/replace"
)

type replace struct {
	args map[string]string
}

// Execute replace cli command
// Replaces old backups from source dir to storage dir
func (r replace) Execute() error {
	sourceDir, err := parseDir(r.args, sourceParam)
	if err != nil {
		return err
	}

	storageDir, err := parseDir(r.args, storageParam)
	if err != nil {
		return err
	}

	daysKeep, err := parseInt(r.args, sourceKeepDaysParam, sourceKeepDaysDefault)
	if err != nil {
		return err
	}

	replacer := service.Replacer{
		SourceDir:  sourceDir,
		StorageDir: storageDir,
		DaysKeep:   daysKeep,
	}

	err = replacer.ReplaceOld()
	if err != nil {
		return err
	}

	return nil
}
