package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func dateOfLastBackup(wslname, backupPath string) (int, error) {
	var numberOfDaysOfLastBackup int

	var latestModTime time.Time
	err := filepath.Walk(backupPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(info.Name(), fmt.Sprintf("%s-backup-", wslname)) {
			if info.ModTime().After(latestModTime) {
				latestModTime = info.ModTime()
			}
		}
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("during the calculation of the date of the last backup: %w", err)
	}

	if !latestModTime.IsZero() {
		numberOfDaysOfLastBackup = int(time.Since(latestModTime).Hours() / 24)
	}

	return numberOfDaysOfLastBackup, nil
}
