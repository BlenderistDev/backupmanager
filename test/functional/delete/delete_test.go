package delete

import (
	"os"
	"testing"
	"time"

	"github.com/BlenderistDev/backupmanager/internal/command"
	"github.com/BlenderistDev/backupmanager/test/tools"
)

const (
	oldName         = "file_old"
	newName         = "file_new"
	oldSingleName   = "single"
	oldSeveralName1 = "several1"
	oldSeveralName2 = "several2"
	oldSeveralName3 = "several3"
)

func TestDelete(t *testing.T) {
	defer func() {
		_ = os.RemoveAll(tools.StorageDir)
	}()

	now := time.Now()
	oldTime := now.Add(-1 * time.Hour * 24 * 8)
	oldTimeSingle, err := time.Parse(time.DateOnly, "2023-07-20")
	if err != nil {
		t.Error(err)
	}
	oldTimeSeveral1, err := time.Parse(time.DateOnly, "2023-07-04")
	if err != nil {
		t.Error(err)
	}
	oldTimeSeveral2, err := time.Parse(time.DateOnly, "2023-07-05")
	if err != nil {
		t.Error(err)
	}
	oldTimeSeveral3, err := time.Parse(time.DateOnly, "2023-07-06")
	if err != nil {
		t.Error(err)
	}

	newPath := tools.GetStoragePath(newName, now)
	oldPath := tools.GetStoragePath(oldName, oldTime)
	singlePath := tools.GetStoragePath(oldSingleName, oldTimeSingle)
	severalPath1 := tools.GetStoragePath(oldSeveralName1, oldTimeSeveral1)
	severalPath2 := tools.GetStoragePath(oldSeveralName2, oldTimeSeveral2)
	severalPath3 := tools.GetStoragePath(oldSeveralName3, oldTimeSeveral3)

	tools.CreateStubFile(t, now, newPath)
	tools.CreateStubFile(t, oldTime, oldPath)
	tools.CreateStubFile(t, oldTimeSingle, singlePath)
	tools.CreateStubFile(t, oldTimeSeveral1, severalPath1)
	tools.CreateStubFile(t, oldTimeSeveral2, severalPath2)
	tools.CreateStubFile(t, oldTimeSeveral3, severalPath3)

	args := []string{"", "delete", "--storage", tools.StorageDir, "--storage-keep", "7"}

	os.Args = args

	err = command.GetCommand().Execute()
	if err != nil {
		t.Error(err)
	}

	tools.CheckFileExist(t, newPath)
	tools.CheckFileExist(t, singlePath)
	tools.CheckFileExist(t, severalPath1)

	tools.CheckFileNotExist(t, severalPath2)
	tools.CheckFileNotExist(t, severalPath3)
}
