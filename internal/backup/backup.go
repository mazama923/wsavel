package backup

import (
	"fmt"
	"time"
)

func BackupWSL(wslName, backupPath string, maxKeep, minDays int, compress bool) (string, error) {
	if err := checkBackupPath(backupPath); err != nil {
		return "", err
	}

	numberOfDaysOfLastBackup, err := dateOfLastBackup(wslName, backupPath)
	if err != nil {
		return "", err
	}

	if numberOfDaysOfLastBackup < minDays {
		return "", nil
	}

	date := time.Now().Format("2006-01-02") // Format YYYY-MM-DD
	backupFilePath := fmt.Sprintf("%s\\%s-backup-%s.tar", backupPath, wslName, date)

	if err := exportWSL(wslName, backupFilePath); err != nil {
		return "", err
	}

	if compress {
		if err := compressBackup(backupFilePath); err != nil {
			return "", err
		}
		backupFilePath += ".gz"
	}

  if err := cleanBackups(wslName, backupFilePath, maxKeep); err != nil {
		return "", err
  }

	return backupFilePath, nil
}
