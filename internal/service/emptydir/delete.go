package emptydir

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// Deleter struct for deleting empty dirs
type Deleter struct {
	StorageDir string
}

// DeleteEmptyDirs delete empty dirs from storage folder
func (d Deleter) DeleteEmptyDirs() error {
	emptyDirs, err := d.getEmptyDirs()
	if err != nil {
		return err
	}

	if len(emptyDirs) > 0 {
		for _, dir := range emptyDirs {
			err = os.Remove(dir)
			if err != nil {
				return err
			}
			log.Println("dir deleted", dir)
		}

		err = d.DeleteEmptyDirs()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d Deleter) getEmptyDirs() ([]string, error) {
	emptyDirs := make([]string, 0)
	if err := filepath.WalkDir(d.StorageDir, func(path string, f fs.DirEntry, err error) error {
		if !f.IsDir() {
			return nil
		}
		dir, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		if len(dir) == 0 {
			emptyDirs = append(emptyDirs, path)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return emptyDirs, nil
}
