package main

import (
	"flag"
	"os"

	"github.com/charmbracelet/log"
	"github.com/mazama923/wsavel/internal/argument"
	"github.com/mazama923/wsavel/internal/backup"
	"github.com/mazama923/wsavel/internal/system"
)

func main() {
	// Check if the OS is Windows before anything else
	if err := system.CheckWindows(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	flag.Parse()

	argumentLoad, err := argument.ParseArgs()
	if err != nil {
		log.Error("While loading arguments", "err", err)
		os.Exit(1)
	}

	backupFilePath, err := backup.BackupWSL(argumentLoad.WSLName, argumentLoad.BackupPath, argumentLoad.MaxKeep, argumentLoad.MinDays, argumentLoad.Compress)
	if err != nil {
		log.Error("While performing backup", "err", err)
		os.Exit(1)
	}

	if backupFilePath != "" {
		log.Info("Backup available: %s", backupFilePath)
	} else {
		log.Info("No need to make a backup", "err", "The last backup is less than %s days old", argumentLoad.MinDays)
	}
}
