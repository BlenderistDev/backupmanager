package emptydir

import (
	"fmt"
	"os"
	"testing"

	"github.com/BlenderistDev/backupmanager/internal/command"
	"github.com/BlenderistDev/backupmanager/test/tools"
)

func TestDeleteEmptyDirs(t *testing.T) {
	defer func() {
		_ = os.RemoveAll(tools.StorageDir)
	}()

	emptyDirPath := fmt.Sprintf("%s/%s", tools.StorageDir, "empty")
	notEmptyDirPath := fmt.Sprintf("%s/%s", tools.StorageDir, "notEmpty")

	emptyOuterDirPath := fmt.Sprintf("%s/%s", tools.StorageDir, "empty_outer")
	emptyInnerDirPath := fmt.Sprintf("%s/%s", emptyOuterDirPath, "empty_inner")

	notEmptyOuterDirPath := fmt.Sprintf("%s/%s", tools.StorageDir, "not_empty_outer")
	notEmptyInnerDirPath := fmt.Sprintf("%s/%s", notEmptyOuterDirPath, "not_empty_inner")

	err := os.MkdirAll(emptyDirPath, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(notEmptyDirPath, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(emptyInnerDirPath, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(notEmptyInnerDirPath, os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	_, err = os.Create(fmt.Sprintf("%s/%s", notEmptyDirPath, "file"))
	if err != nil {
		t.Error(err)
	}
	_, err = os.Create(fmt.Sprintf("%s/%s", notEmptyInnerDirPath, "file"))
	if err != nil {
		t.Error(err)
	}

	args := []string{"", "emptydir", "--storage", tools.StorageDir}

	os.Args = args

	cmd, err := command.GetCommand()
	if err != nil {
		t.Error(err)
	}
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	tools.CheckFileExist(t, notEmptyDirPath)
	tools.CheckFileExist(t, notEmptyInnerDirPath)
	tools.CheckFileExist(t, notEmptyOuterDirPath)

	tools.CheckFileNotExist(t, emptyDirPath)
	tools.CheckFileNotExist(t, emptyOuterDirPath)
	tools.CheckFileNotExist(t, emptyInnerDirPath)
}
