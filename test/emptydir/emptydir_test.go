package emptydir

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/BlenderistDev/backupmanager/internal/command"
)

func TestDeleter_DeleteEmptyDirs(t *testing.T) {
	const storageDir = "test_replacer_storage_dir"
	defer func() {
		_ = os.RemoveAll(storageDir)
	}()

	emptyDirPath := fmt.Sprintf("%s/%s", storageDir, "empty")
	notEmptyDirPath := fmt.Sprintf("%s/%s", storageDir, "notEmpty")

	emptyOuterDirPath := fmt.Sprintf("%s/%s", storageDir, "empty_outer")
	emptyInnerDirPath := fmt.Sprintf("%s/%s", emptyOuterDirPath, "empty_inner")

	notEmptyOuterDirPath := fmt.Sprintf("%s/%s", storageDir, "not_empty_outer")
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

	args := []string{"", "emptydir", "--storage", storageDir}

	os.Args = args

	cmd, err := command.GetCommand()
	if err != nil {
		t.Error(err)
	}
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(notEmptyDirPath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(notEmptyInnerDirPath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(notEmptyOuterDirPath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(emptyDirPath); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("dir " + emptyDirPath + " have not been deleted")
		}
	}

	if _, err := os.Stat(emptyOuterDirPath); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("dir " + emptyOuterDirPath + " have not been deleted")
		}
	}

	if _, err := os.Stat(emptyInnerDirPath); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("dir " + emptyInnerDirPath + " have not been deleted")
		}
	}

}
