package argument

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/charmbracelet/log"
)

type Config struct {
	WSLName    string
	BackupPath string
	MaxKeep int
	MinDays    int
	Compress   bool
}

func ParseArgs() (*Config, error) {
	// Check with only 4 parameters because <compress> can be empty
	if len(flag.Args()) < 4 {
		return nil, errors.New("usage: wsavel <wsl_name> <backup_path> <max_keep> <min_days> <compress>")
	}

	maxKeep, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		return nil, fmt.Errorf("invalid <max_keep>: %s (must be a number)", flag.Arg(2))
	}

	minDays, err := strconv.Atoi(flag.Arg(3))
	if err != nil {
		return nil, fmt.Errorf("invalid <min_days>: %s (must be a number)", flag.Arg(3))
	}

	var compress bool

	switch flag.Arg(4) {
	case "":
		log.Warn("No value for <compress> by default there will be no compression")
		compress = false
	case "true":
		compress = true
	case "false":
		compress = false
	default:
		return nil, fmt.Errorf("invalid value for <compress> must be true or false")
	}

	return &Config{
		WSLName:    flag.Arg(0),
		BackupPath: flag.Arg(1),
		MaxKeep: maxKeep,
		MinDays:    minDays,
		Compress:   compress,
	}, nil
}
