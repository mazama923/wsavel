package argument

import (
	"flag"
	"testing"
)

// Test with valid arguments
func TestParseArgs_Valid(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.Parse([]string{"Ubuntu-20.04", "/backups", "5", "14", "true"})

	cfg, err := ParseArgs()
	if err != nil {
		t.Fatalf("ParseArgs failed: %v", err)
	}

	// Verify expected values
	if cfg.WSLName != "Ubuntu-20.04" {
		t.Errorf("Expected WSLName 'Ubuntu-20.04', got '%s'", cfg.WSLName)
	}
	if cfg.BackupPath != "/backups" {
		t.Errorf("Expected BackupPath '/backups', got '%s'", cfg.BackupPath)
	}
	if cfg.MaxKeep != 5 {
		t.Errorf("Expected MaxKeep 5, got %d", cfg.MaxKeep)
	}
	if cfg.MinDays != 14 {
		t.Errorf("Expected MinDays 14, got %d", cfg.MinDays)
	}
	if cfg.Compress != true {
		t.Errorf("Expected Compress true, got %v", cfg.Compress)
	}
}

// Test invalid <max_keep> (should be an integer)
func TestParseArgs_InvalidMaxKeep(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.Parse([]string{"Ubuntu-20.04", "/backups", "abc", "14", "true"})

	_, err := ParseArgs()
	if err == nil {
		t.Fatal("Expected error for invalid <max_keep>, but got nil")
	}
}

// Test invalid <min_days> (should be an integer)
func TestParseArgs_InvalidMinDays(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.Parse([]string{"Ubuntu-20.04", "/backups", "5", "xyz", "true"})

	_, err := ParseArgs()
	if err == nil {
		t.Fatal("Expected error for invalid <min_days>, but got nil")
	}
}

// Test invalid <compress> value (should be "true" or "false")
func TestParseArgs_InvalidCompress(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.Parse([]string{"Ubuntu-20.04", "/backups", "5", "14", "maybe"})

	_, err := ParseArgs()
	if err == nil {
		t.Fatal("Expected error for invalid <compress>, but got nil")
	}
}

// Test missing <compress> (should default to false with a warning)
func TestParseArgs_MissingCompress(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.Parse([]string{"Ubuntu-20.04", "/backups", "5", "14"})

	cfg, err := ParseArgs()
	if err != nil {
		t.Fatalf("ParseArgs failed: %v", err)
	}

	// Ensure default value is false when <compress> is missing
	if cfg.Compress != false {
		t.Errorf("Expected Compress false when missing, got %v", cfg.Compress)
	}
}

// Test when not enough arguments are provided (should return an error)
func TestParseArgs_NotEnoughArguments(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.Parse([]string{"Ubuntu-20.04", "/backups"})

	_, err := ParseArgs()
	if err == nil {
		t.Fatal("Expected error for missing arguments, but got nil")
	}
}
