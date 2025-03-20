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
	backupFileName := fmt.Sprintf("%s-backup-%s.tar", wslName, date)
	backupFilePath := fmt.Sprintf("%s\\%s", backupPath, backupFileName)

	if err := exportWSL(wslName, backupFilePath); err != nil {
		return "", err
	}

	if compress {
		if err := compressBackup(wslName, backupPath, backupFileName); err != nil {
			return "", err
		}
		backupFilePath += ".gz"
	}

	if err := cleanBackups(wslName, backupPath, maxKeep); err != nil {
		return "", err
	}

	return backupFilePath, nil
}
