package backup

import (
	"fmt"
	"os/exec"

	"github.com/mazama923/wsavel/internal/ui"
)

func compressBackup(backupFilePath string) error {
	go ui.StartSpinner("Compressing backup...")
	defer ui.StopSpinner()

	cmd := exec.Command("wsl.exe", "gzip", backupFilePath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("compressing backup file: %w", err)
	}
	return nil
}
