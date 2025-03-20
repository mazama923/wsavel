package backup

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mazama923/wsavel/internal/ui"
)

func compressBackup(wslName, backupPath, backupFileName string) error {
	// Convert Windows path to WSL path
	wslBackupPath := strings.Replace(backupPath, "C:\\", "/mnt/c/", 1)
	wslBackupPath = strings.Replace(wslBackupPath, "\\", "/", -1)

	cmd := exec.Command("wsl", "-d", wslName, "sh", "-c", fmt.Sprintf("gzip %s/%s", wslBackupPath, backupFileName))
	logExecCMD := fmt.Sprintf("Compressing backup.: %s", cmd.String())
	ui.StartSpinner(logExecCMD)
	defer ui.StopSpinner()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("compressing backup: %w", err)
	}
	return nil
}
