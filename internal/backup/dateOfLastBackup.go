package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func dateOfLastBackup(wslName, backupPath string) (int, error) {
	numberOfDaysOfLastBackup := 9999 // Default value if no file exists to trigger numberOfDaysOfLastBackup < mindays
	var latestModTime time.Time
	found := false

	err := filepath.Walk(backupPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(info.Name(), fmt.Sprintf("%s-backup-", wslName)) {
			if !found || info.ModTime().After(latestModTime) {
				latestModTime = info.ModTime()
				found = true
			}
		}
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("error while checking the last backup date: %w", err)
	}

	if found {
		numberOfDaysOfLastBackup = int(time.Since(latestModTime).Hours() / 24)
	}

	return numberOfDaysOfLastBackup, nil
}
