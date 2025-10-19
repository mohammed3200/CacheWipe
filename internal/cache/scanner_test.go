package cache

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewScanner(t *testing.T) {
	scanner := NewScanner()
	if scanner == nil {
		t.Fatal("Expected scanner to be non-nil")
	}
	if scanner.IsStopped() {
		t.Fatal("New scanner should not be stopped")
	}
}

func TestScannerStop(t *testing.T) {
	scanner := NewScanner()
	scanner.Stop()
	if !scanner.IsStopped() {
		t.Fatal("Scanner should be stopped after calling Stop()")
	}
}

func TestCalculateDirSize(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test files
	testFile := filepath.Join(tmpDir, "test.txt")
	content := []byte("test content")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	size, count, err := calculateDirSize(tmpDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedSize := int64(len(content))
	if size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	if count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}
}

func TestScanEmptyDirectory(t *testing.T) {
	scanner := NewScanner()
	result, err := scanner.Scan(func(progress int64) {})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected scan result to be non-nil")
	}
}

func TestGetCachePaths(t *testing.T) {
	paths := getCachePaths()
	if paths == nil || len(paths) == 0 {
		t.Fatal("Expected cache paths to be non-empty")
	}
}

func BenchmarkCalculateDirSize(b *testing.B) {
	tmpDir := b.TempDir()

	for i := 0; i < 100; i++ {
		testFile := filepath.Join(tmpDir, "test_"+string(rune(i))+".txt")
		os.WriteFile(testFile, []byte("test"), 0644)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateDirSize(tmpDir)
	}
}
