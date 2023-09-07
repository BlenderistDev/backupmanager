package tools

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// SourceDir test source directory
const SourceDir = "test_replacer_source_dir"

// StorageDir test storage directory
const StorageDir = "test_replacer_storage_dir"

// CheckFileExist checks file exist
func CheckFileExist(t *testing.T, path string) {
	if _, err := os.Stat(path); err != nil {
		t.Error(err)
	}
}

// CheckFileNotExist checks file not exist
func CheckFileNotExist(t *testing.T, path string) {
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("file " + path + " have not been moved")
		}
	}
}

// CreateStubFile creates stub file with setting chtimes
func CreateStubFile(t *testing.T, time time.Time, path string) {
	var dir string
	i := strings.LastIndex(path, "/")
	if i > 0 {
		dir = path[:i]
	}
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	_, err = os.Create(path)
	if err != nil {
		t.Error(err)
	}

	err = os.Chtimes(path, time, time)
	if err != nil {
		t.Error(err)
	}
}

// GetSourcePath returns source path for file
func GetSourcePath(name string) string {
	return SourceDir + "/" + name
}

// GetStoragePath returns test storage path
func GetStoragePath(name string, oldTime time.Time) string {
	return fmt.Sprintf("%s/%d/%s/%s/%s", StorageDir, oldTime.Year(), oldTime.Month(), oldTime.Format(time.DateOnly), name)
}
