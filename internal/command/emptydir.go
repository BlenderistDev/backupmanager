package command

import (
	service "github.com/BlenderistDev/backupmanager/internal/service/emptydir"
)

type emptydir struct {
	args map[string]string
}

// Execute emptydir cli command
// Deletes empty directories from directory
func (e emptydir) Execute() error {
	storageDir, err := parseDir(e.args, storageParam)
	if err != nil {
		return err
	}

	deleter := service.Deleter{StorageDir: storageDir}

	err = deleter.DeleteEmptyDirs()
	if err != nil {
		return err
	}

	return nil
}
