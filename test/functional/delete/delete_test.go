package delete

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/BlenderistDev/backupmanager/internal/command"
)

const storageDir = "test_replacer_storage_dir"

func TestDeleter_DeleteOld(t *testing.T) {

	defer func() {
		_ = os.RemoveAll(storageDir)
	}()

	now := time.Now()

	oldPathSingle, err := createStubFile(now.Add(-1*time.Hour*24*8), "%s/%d/%s/%s")
	if err != nil {
		t.Error(err)
	}

	oldPathSeveral1, err := createStubFile(now.Add(-1*time.Hour*24*16), "%s/%d/%s/%s-1")
	if err != nil {
		t.Error(err)
	}

	oldPathSeveral2, err := createStubFile(now.Add(-1*time.Hour*24*16), "%s/%d/%s/%s-2")
	if err != nil {
		t.Error(err)
	}

	newPath, err := createStubFile(now, "%s/%d/%s/%s")
	if err != nil {
		t.Error(err)
	}

	args := []string{"", "delete", "--storage", storageDir, "--storage-keep", "7"}

	os.Args = args

	cmd, err := command.GetCommand()
	if err != nil {
		t.Error(err)
	}
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(newPath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(oldPathSingle); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(oldPathSeveral1); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(oldPathSeveral2); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("file " + oldPathSeveral2 + " have not been deleted")
		}
	}
}

func createStubFile(t time.Time, format string) (string, error) {
	dir := fmt.Sprintf(format, storageDir, t.Year(), t.Month(), t.Format(time.DateOnly))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	path := dir + "/file"
	_, err := os.Create(path)
	if err != nil {
		return "", err
	}

	err = os.Chtimes(path, t, t)
	if err != nil {
		return "", err
	}

	return path, nil
}
