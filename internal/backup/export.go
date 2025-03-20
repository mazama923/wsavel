package backup

import (
	"fmt"
	"os/exec"

	"github.com/mazama923/wsavel/internal/ui"
)

func exportWSL(wslName, backupFilePath string) error {
	cmd := exec.Command("wsl", "--export", wslName, backupFilePath)
	logExecCMD := fmt.Sprintf("Exporting WSL: %s", cmd.String())
	ui.StartSpinner(logExecCMD)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exporting WSL: %w", err)
	}
	return nil
}
