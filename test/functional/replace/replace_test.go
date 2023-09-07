package replace

import (
	"os"
	"testing"
	"time"

	"github.com/BlenderistDev/backupmanager/internal/command"
	"github.com/BlenderistDev/backupmanager/test/tools"
)

const (
	oldName = "/file_old"
	newName = "/file_new"
)

func TestReplaceOld(t *testing.T) {
	defer func() {
		_ = os.RemoveAll(tools.SourceDir)
		_ = os.RemoveAll(tools.StorageDir)
	}()

	now := time.Now()
	oldTime := now.Add(-1 * time.Hour * 24 * 8)

	newSourcePath := tools.GetSourcePath(newName)
	oldSourcePath := tools.GetSourcePath(oldName)
	oldStoragePath := tools.GetStoragePath(oldName, oldTime)

	tools.CreateStubFile(t, now, newSourcePath)
	tools.CreateStubFile(t, oldTime, oldSourcePath)

	args := []string{"", "replace", "--source", tools.SourceDir, "--storage", tools.StorageDir, "--source-keep", "7"}

	os.Args = args

	cmd, err := command.GetCommand()
	if err != nil {
		t.Error(err)
	}
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	tools.CheckFileExist(t, newSourcePath)
	tools.CheckFileNotExist(t, oldSourcePath)
	tools.CheckFileExist(t, oldStoragePath)
}
