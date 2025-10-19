package models

import (
	"testing"
	"time"
)

func TestSystemStats(t *testing.T) {
	stats := &SystemStats{
		TotalDiskSpace: 1000000000,
		UsedDiskSpace:  500000000,
		FreeDiskSpace:  500000000,
		CacheSize:      100000000,
	}

	if stats.TotalDiskSpace != 1000000000 {
		t.Error("SystemStats TotalDiskSpace mismatch")
	}
}

func TestCleanupStats(t *testing.T) {
	now := time.Now()
	stats := &CleanupStats{
		StartTime:    now,
		EndTime:      now.Add(time.Second * 5),
		Duration:     time.Second * 5,
		DeletedSize:  100000,
		DeletedCount: 50,
		ErrorCount:   2,
		Status:       "Completed",
		Success:      true,
	}

	if !stats.Success {
		t.Error("Expected Success to be true")
	}

	if stats.DeletedCount != 50 {
		t.Error("DeletedCount mismatch")
	}
}

func TestConfig(t *testing.T) {
	config := &Config{
		AutoCleanup:     true,
		CleanupInterval: "weekly",
		LogCleanup:      true,
		DryRun:          false,
	}

	if !config.AutoCleanup {
		t.Error("Expected AutoCleanup to be true")
	}

	if config.CleanupInterval != "weekly" {
		t.Error("CleanupInterval mismatch")
	}
}
