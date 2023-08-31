package replace

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Replacer struct {
	SourceDir  string
	StorageDir string
	DaysKeep   int
}

type fileToReplace struct {
	date time.Time
	name string
}

func (r Replacer) ReplaceOld() error {
	files, err := r.getFilesToReplace()
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	err = r.replaceFiles(files)
	if err != nil {
		return err
	}

	return nil
}

func (r Replacer) getFilesToReplace() ([]fileToReplace, error) {
	dirData, err := os.ReadDir(r.SourceDir)
	if err != nil {
		return nil, err
	}

	toReplace := make([]fileToReplace, 0)
	for _, v := range dirData {
		if v.IsDir() {
			continue
		}

		fileInfo, err := v.Info()
		if err != nil {
			return nil, err
		}

		if time.Now().Before(fileInfo.ModTime().Add(time.Duration(r.DaysKeep) * time.Hour * 24)) {
			continue
		}

		toReplace = append(toReplace, fileToReplace{
			date: fileInfo.ModTime(),
			name: v.Name(),
		})
	}

	return toReplace, nil
}

func (r Replacer) replaceFiles(toReplace []fileToReplace) error {
	for _, f := range toReplace {
		year := f.date.Year()
		month := f.date.Month().String()
		date := f.date.Format(time.DateOnly)

		newDir := fmt.Sprintf("%s/%d/%s/%s", r.StorageDir, year, month, date)

		err := os.MkdirAll(newDir, os.ModePerm)
		if err != nil {
			return err
		}

		from, to := fmt.Sprintf("%s/%s", r.SourceDir, f.name), fmt.Sprintf("%s/%s", newDir, f.name)
		err = os.Rename(from, to)
		if err != nil {
			return err
		}
		log.Println(fmt.Sprintf("file %s replaced to %s", from, to))
	}

	return nil
}
