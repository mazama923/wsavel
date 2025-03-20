package backup

import (
	"fmt"
	"os"
)

func checkBackupPath(backupPath string) error {
	if err := os.MkdirAll(backupPath, os.ModePerm); err != nil {
		return fmt.Errorf("during the creation of the backup folder: %w", err)
	}
	return nil
}
