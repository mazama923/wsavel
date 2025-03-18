package backup

import (
	"fmt"
	"os/exec"

	"github.com/mazama923/wsavel/internal/ui"
)

func compressBackup(wslName, backupPath, backupFileName string) error {
	go ui.StartSpinner("Compressing backup...")
	defer ui.StopSpinner()

	cmd := exec.Command("wsl", "-d", wslName, "-- cd ", backupPath, "gzip", backupFileName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("compressing backup: %w", err)
	}
	return nil
}
