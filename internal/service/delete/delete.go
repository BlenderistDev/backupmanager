package delete

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Deleter struct for deleting old files from storage
type Deleter struct {
	StorageDir string
	DaysKeep   int
}

// DeleteOld deletes old files from storage.
// Keep one file from every week
func (d Deleter) DeleteOld() error {
	oldList := make(map[string][]string)
	if err := filepath.WalkDir(d.StorageDir, func(path string, f fs.DirEntry, err error) error {
		if f.IsDir() {
			return nil
		}

		info, err := f.Info()
		if err != nil {
			return err
		}

		if info.ModTime().Add(time.Duration(d.DaysKeep) * time.Hour * 24).Before(time.Now()) {
			year, week := info.ModTime().ISOWeek()
			key := fmt.Sprintf("%d-%d", year, week)
			if _, ok := oldList[key]; ok {
				oldList[key] = append(oldList[key], path)
			} else {
				oldList[key] = []string{path}
			}
		}

		return nil
	}); err != nil {
		return err
	}

	for _, files := range oldList {
		for i, file := range files {
			if i == 0 {
				continue
			}
			err := os.Remove(file)
			if err != nil {
				return err
			}
			log.Println("file deleted", file)
		}
	}

	return nil
}
