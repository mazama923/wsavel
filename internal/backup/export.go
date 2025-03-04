package backup

import (
	"fmt"
	"os/exec"

	"github.com/mazama923/wsavel/internal/ui"
)

func exportWSL(wslname, backupFilePath string) error {
	go ui.StartSpinner("Exporting WSL...")
	defer ui.StopSpinner()

	cmd := exec.Command("wsl", "--export", wslname, backupFilePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exporting WSL: %w", err)
	}
	return nil
}
