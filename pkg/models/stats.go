package models

import "time"

type SystemStats struct {
	TotalDiskSpace int64
	UsedDiskSpace  int64
	FreeDiskSpace  int64
	CacheSize      int64
	OptimizedSize  int64
}

type CleanupStats struct {
	StartTime    time.Time
	EndTime      time.Time
	Duration     time.Duration
	DeletedSize  int64
	DeletedCount int64
	ErrorCount   int
	Status       string
	Success      bool
}

type Config struct {
	AutoCleanup     bool
	CleanupInterval string
	SelectedPaths   []string
	LogCleanup      bool
	DryRun          bool
}

type CleanupLog struct {
	Timestamp    time.Time
	DeletedSize  int64
	DeletedCount int64
	Status       string
}
