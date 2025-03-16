package main

import (
	"fmt"
	"time"

	"github.com/mazama923/wsavel/internal/ui"
)

func main() {
	go ui.StartSpinner("Exporting WSL...")
	fmt.Println("cc")
	// add wait 8sec
	time.Sleep(8 * time.Second)
	defer ui.StopSpinner()
}
