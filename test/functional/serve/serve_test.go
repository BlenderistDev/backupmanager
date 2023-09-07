package serve

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/BlenderistDev/backupmanager/internal/command"
)

const sourceDir = "test_replacer_source_dir"
const storageDir = "test_replacer_storage_dir"
const oldName = "file_old"
const newName = "file_new"
const oldSingleName = "single"
const oldSeveralName1 = "several1"
const oldSeveralName2 = "several2"
const oldSeveralName3 = "several3"

func TestServe(t *testing.T) {

	defer func() {
		_ = os.RemoveAll(sourceDir)
		_ = os.RemoveAll(storageDir)
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

	// prepare test directories and files
	err = os.Mkdir(sourceDir, os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	createStubFile(t, now, newName)
	createStubFile(t, oldTime, oldName)
	createStubFile(t, oldTimeSingle, oldSingleName)
	createStubFile(t, oldTimeSeveral1, oldSeveralName1)
	createStubFile(t, oldTimeSeveral2, oldSeveralName2)
	createStubFile(t, oldTimeSeveral3, oldSeveralName3)

	args := []string{"", "serve", "--source", sourceDir, "--storage", storageDir, "--source-keep", "7", "--storage-keep", "14"}

	os.Args = args

	cmd, err := command.GetCommand()
	if err != nil {
		t.Error(err)
	}
	go func() {
		err = cmd.Execute()
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	checkFileExist(t, getSourcePath(newName))
	checkFileExist(t, getStoragePath(oldName, oldTime))
	checkFileExist(t, getStoragePath(oldSingleName, oldTimeSingle))
	checkFileExist(t, getStoragePath(oldSeveralName1, oldTimeSeveral1))

	checkFileNotExist(t, getSourcePath(oldName))
	checkFileNotExist(t, getSourcePath(oldSingleName))
	checkFileNotExist(t, getSourcePath(oldSeveralName1))
	checkFileNotExist(t, getSourcePath(oldSeveralName2))
	checkFileNotExist(t, getSourcePath(oldSeveralName3))

	checkFileNotExist(t, getStoragePath(oldSeveralName2, oldTimeSeveral2))
	checkFileNotExist(t, getStoragePath(oldSeveralName3, oldTimeSeveral3))
}

func getStoragePath(name string, oldTime time.Time) string {
	return fmt.Sprintf("%s/%d/%s/%s/%s", storageDir, oldTime.Year(), oldTime.Month(), oldTime.Format(time.DateOnly), name)
}

func checkFileExist(t *testing.T, path string) {
	if _, err := os.Stat(path); err != nil {
		t.Error(err)
	}
}

func checkFileNotExist(t *testing.T, path string) {
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("file " + path + " have not been moved")
		}
	}
}

func createStubFile(t *testing.T, time time.Time, name string) {
	path := getSourcePath(name)
	_, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}

	err = os.Chtimes(path, time, time)
	if err != nil {
		t.Error(err)
	}
}

func getSourcePath(name string) string {
	return sourceDir + "/" + name
}
