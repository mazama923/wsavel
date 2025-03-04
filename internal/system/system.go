package system

import (
	"errors"
	"runtime"
)

// Check if we are on a Windows system because when we export a WSL it is shutdown
func CheckWindows() error {
	if runtime.GOOS != "windows" {
		return errors.New("this program can only be run on Windows")
	}
	return nil
}
