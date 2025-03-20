package backup

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mazama923/wsavel/internal/ui"
)

func cleanBackups(wslName, backupPath string, maxKeep int) error {
	ui.UpdateSpinnerMessage("Cleaning backups ...")

	files, err := os.ReadDir(backupPath)
	if err != nil {
		return err
	}

	var fileInfos []os.FileInfo
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), wslName+"-backup") || strings.Contains(info.Name(), wslName) {
			fileInfos = append(fileInfos, info)
		}
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
	})

	for i := maxKeep; i < len(fileInfos); i++ {
		err := os.Remove(filepath.Join(backupPath, fileInfos[i].Name()))
		if err != nil {
			return err
		}
	}

	return nil
}
