package cache

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewCleaner(t *testing.T) {
	cleaner := NewCleaner()
	if cleaner == nil {
		t.Fatal("Expected cleaner to be non-nil")
	}
}

func TestCleanerStop(t *testing.T) {
	cleaner := NewCleaner()
	cleaner.Stop()
	if !cleaner.IsStopped() {
		t.Fatal("Cleaner should be stopped")
	}
}

func TestCleanPath(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	content := []byte("test content to delete")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	cleaner := NewCleaner()
	deletedSize, _, err := cleaner.cleanPath(tmpDir)

	if err != nil && err != os.ErrNotExist {
		t.Fatalf("Expected no error or NotExist, got %v", err)
	}

	if deletedSize < int64(len(content)) {
		t.Errorf("Expected deleted size >= %d, got %d", len(content), deletedSize)
	}
}

func TestDryRun(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	items := []CacheItem{
		{
			Path:      tmpDir,
			Size:      1024,
			FileCount: 1,
			Selected:  true,
		},
		{
			Path:      tmpDir,
			Size:      2048,
			FileCount: 2,
			Selected:  false,
		},
	}

	cleaner := NewCleaner()
	totalSize, totalCount := cleaner.DryRun(items)

	expectedSize := int64(1024)
	expectedCount := int64(1)

	if totalSize != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, totalSize)
	}

	if totalCount != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, totalCount)
	}
}

func TestCleanupWithProgress(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	item := CacheItem{
		Path:      tmpDir,
		Size:      4,
		FileCount: 1,
		Selected:  true,
	}

	progressCalls := 0
	cleaner := NewCleaner()
	result := cleaner.Cleanup([]CacheItem{item}, func(progress CleanupProgress) {
		progressCalls++
	})

	if result == nil {
		t.Fatal("Expected cleanup result to be non-nil")
	}

	if progressCalls == 0 {
		t.Error("Expected progress callback to be called")
	}
}
