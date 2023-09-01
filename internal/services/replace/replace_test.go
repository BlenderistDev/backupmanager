package replace

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestReplacer_ReplaceOld(t *testing.T) {
	const sourceDir = "test_replacer_source_dir"
	const storageDir = "test_replacer_storage_dir"
	const oldSourcePath = sourceDir + "/file_old"
	const newSourcePath = sourceDir + "/file_new"

	defer func() {
		_ = os.RemoveAll(sourceDir)
		_ = os.RemoveAll(storageDir)
	}()

	now := time.Now()
	oldTime := now.Add(-1 * time.Hour * 24 * 8)
	oldResultPath := fmt.Sprintf("%s/%d/%s/%s", storageDir, oldTime.Year(), oldTime.Month(), oldTime.Format(time.DateOnly))

	// prepare test directories and files
	err := os.Mkdir(sourceDir, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	_, err = os.Create(oldSourcePath)
	if err != nil {
		t.Error(err)
	}
	_, err = os.Create(newSourcePath)
	if err != nil {
		t.Error(err)
	}
	err = os.Chtimes(oldSourcePath, oldTime, oldTime)
	if err != nil {
		t.Error(err)
	}

	replacer := Replacer{
		SourceDir:  sourceDir,
		StorageDir: storageDir,
		DaysKeep:   7,
	}

	err = replacer.ReplaceOld()
	if err != nil {
		t.Error(err)
	}

	// check old file is replaced, new is not replaced
	if _, err := os.Stat(newSourcePath); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(oldSourcePath); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("file " + oldSourcePath + " have not been moved")
		}
	}

	if _, err := os.Stat(oldResultPath); err != nil {
		t.Error(err)
	}
}
