package command

import (
	"fmt"
	"strconv"
)

const (
	sourceParam          = "--source"
	storageParam         = "--storage"
	sourceKeepDaysParam  = "--source-keep"
	storageKeepDaysParam = "--storage-keep"
	sleepParam           = "--sleep"

	sourceKeepDaysDefault  = 14
	storageKeepDaysDefault = 30
	sleepDefault           = 1
)

func parseInt(args map[string]string, param string, def int) (int, error) {
	daysKeep := def
	if val, ok := args[param]; ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("error parsing %s: %v", param, err)
		}
		if i <= 0 {
			return 0, fmt.Errorf("%s must be positive integer", param)
		}
		daysKeep = i
	}
	return daysKeep, nil
}

func parseDir(args map[string]string, param string) (string, error) {
	storageDir, ok := args[param]
	if !ok || storageDir == "" {
		return "", fmt.Errorf("no %s param dir provided, provide storage dir with %s", param, param)
	}
	return storageDir, nil
}
