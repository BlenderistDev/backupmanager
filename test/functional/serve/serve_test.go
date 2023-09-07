package serve

import (
	"os"
	"syscall"
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

func TestServe(t *testing.T) {
	defer func() {
		_ = os.RemoveAll(tools.SourceDir)
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

	tools.CreateStubFile(t, now, tools.GetSourcePath(newName))
	tools.CreateStubFile(t, oldTime, tools.GetSourcePath(oldName))
	tools.CreateStubFile(t, oldTimeSingle, tools.GetSourcePath(oldSingleName))
	tools.CreateStubFile(t, oldTimeSeveral1, tools.GetSourcePath(oldSeveralName1))
	tools.CreateStubFile(t, oldTimeSeveral2, tools.GetSourcePath(oldSeveralName2))
	tools.CreateStubFile(t, oldTimeSeveral3, tools.GetSourcePath(oldSeveralName3))

	args := []string{"", "serve", "--source", tools.SourceDir, "--storage", tools.StorageDir, "--source-keep", "7", "--storage-keep", "14"}
	os.Args = args
	go func() {
		err = command.GetCommand().Execute()
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	tools.CheckFileExist(t, tools.GetSourcePath(newName))
	tools.CheckFileExist(t, tools.GetStoragePath(oldName, oldTime))
	tools.CheckFileExist(t, tools.GetStoragePath(oldSingleName, oldTimeSingle))
	tools.CheckFileExist(t, tools.GetStoragePath(oldSeveralName1, oldTimeSeveral1))

	tools.CheckFileNotExist(t, tools.GetSourcePath(oldName))
	tools.CheckFileNotExist(t, tools.GetSourcePath(oldSingleName))
	tools.CheckFileNotExist(t, tools.GetSourcePath(oldSeveralName1))
	tools.CheckFileNotExist(t, tools.GetSourcePath(oldSeveralName2))
	tools.CheckFileNotExist(t, tools.GetSourcePath(oldSeveralName3))

	tools.CheckFileNotExist(t, tools.GetStoragePath(oldSeveralName2, oldTimeSeveral2))
	tools.CheckFileNotExist(t, tools.GetStoragePath(oldSeveralName3, oldTimeSeveral3))
}
