package backup

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/mazama923/wsavel/internal/ui"
)

func compressBackup(backupFilePath string) error {
	go ui.StartSpinner("Compressing backup...")
	defer ui.StopSpinner()

	// Ouvrir le fichier source
	srcFile, err := os.Open(backupFilePath)
	if err != nil {
		return fmt.Errorf("opening backup file: %w", err)
	}
	defer srcFile.Close()

	// Créer le fichier compressé avec l'extension .gz
	destFilePath := backupFilePath + ".gz"
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return fmt.Errorf("creating compressed file: %w", err)
	}
	defer destFile.Close()

	// Créer un writer gzip
	gzipWriter := gzip.NewWriter(destFile)
	defer gzipWriter.Close()

	// Copier les données du fichier source vers le writer gzip
	_, err = io.Copy(gzipWriter, srcFile)
	if err != nil {
		return fmt.Errorf("compressing backup file: %w", err)
	}

	// Fermeture du writer gzip pour s'assurer que tout est bien écrit
	if err := gzipWriter.Close(); err != nil {
		return fmt.Errorf("closing gzip writer: %w", err)
	}

	// Supprimer le fichier original après compression
	if err := os.Remove(backupFilePath); err != nil {
		return fmt.Errorf("deleting original backup file: %w", err)
	}

	return nil
}
