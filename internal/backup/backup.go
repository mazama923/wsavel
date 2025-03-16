package backup

import (
	"fmt"
	"time"
)

func BackupWSL(wslname, backupPath string, maxkeep, mindays int, compress bool) (string, error) {
	if err := checkBackupPath(backupPath); err != nil {
		return "", err
	}

  fmt.Println("Check backup path passed")

	numberOfDaysOfLastBackup, err := dateOfLastBackup(wslname, backupPath)
	if err != nil {
		return "", err
	}

  fmt.Println("Check date of last backup passed")

	if numberOfDaysOfLastBackup < mindays {
		return "", nil
	}

  fmt.Println("Check min days passed")

	date := time.Now().Format("2006-01-02") // Format YYYY-MM-DD
	backupFilePath := fmt.Sprintf("%s/%s-backup-%s.tar", backupPath, wslname, date)

	if err := exportWSL(wslname, backupFilePath); err != nil {
		return "", err
	}
  fmt.Println("Export WSL passed")

	if compress {
		if err := compressBackup(backupFilePath); err != nil {
			return "", err
		}
		backupFilePath += ".gz"
	}

	return backupFilePath, nil
}
