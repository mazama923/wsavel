package backup

import (
	"fmt"
	"time"
)

func BackupWSL(wslname, backupPath string, maxkeep, mindays int, compress bool) (string, error) {
	if err := checkBackupPath(backupPath); err != nil {
		return "", err
	}

	numberOfDaysOfLastBackup, err := dateOfLastBackup(wslname, backupPath)
	if err != nil {
		return "", err
	}

	if numberOfDaysOfLastBackup < mindays {
		return "", nil
	}

	date := time.Now().Format("2006-01-02") // Format YYYY-MM-DD
	backupFilePath := fmt.Sprintf("%s/%s-backup-%s.tar", backupPath, wslname, date)

	if err := exportWSL(wslname, backupFilePath); err != nil {
		return "", err
	}

	if compress {
		if err := compressBackup(backupFilePath); err != nil {
			return "", err
		}
	}

	return backupFilePath, nil
}
