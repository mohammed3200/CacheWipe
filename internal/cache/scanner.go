package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
)

type ScanResult struct {
	TotalSize  int64
	FileCount  int64
	Items      map[string]CacheItem
	Errors     []error
	Categories map[string]CategoryStats
}

type CacheItem struct {
	Path      string
	Size      int64
	FileCount int64
	Category  string
	Selected  bool
}

type CategoryStats struct {
	Name      string
	Size      int64
	FileCount int64
	Paths     []string
}

type Scanner struct {
	stopped int32
	mu      sync.Mutex
}

func NewScanner() *Scanner {
	return &Scanner{
		stopped: 0,
	}
}

func (s *Scanner) Stop() {
	atomic.StoreInt32(&s.stopped, 1)
}

func (s *Scanner) IsStopped() bool {
	return atomic.LoadInt32(&s.stopped) == 1
}

func (s *Scanner) Scan(onProgress func(int64)) (*ScanResult, error) {
	result := &ScanResult{
		Items:      make(map[string]CacheItem),
		Categories: make(map[string]CategoryStats),
		Errors:     []error{},
	}

	atomic.StoreInt32(&s.stopped, 0)
	cachePaths := getCachePaths()
	totalPaths := 0

	for _, paths := range cachePaths {
		totalPaths += len(paths)
	}

	scannedPaths := 0

	for category, paths := range cachePaths {
		categoryStats := CategoryStats{
			Name:  category,
			Paths: []string{},
		}

		for _, path := range paths {
			if s.IsStopped() {
				break
			}

			if _, err := os.Stat(path); os.IsNotExist(err) {
				scannedPaths++
				onProgress(int64(scannedPaths * 100 / totalPaths))
				continue
			}

			size, count, err := calculateDirSize(path)
			if err != nil {
				result.Errors = append(result.Errors, fmt.Errorf("error scanning %s: %w", path, err))
				scannedPaths++
				onProgress(int64(scannedPaths * 100 / totalPaths))
				continue
			}

			if size > 0 {
				result.TotalSize += size
				result.FileCount += count

				item := CacheItem{
					Path:      path,
					Size:      size,
					FileCount: count,
					Category:  category,
					Selected:  true,
				}
				result.Items[path] = item

				categoryStats.Size += size
				categoryStats.FileCount += count
				categoryStats.Paths = append(categoryStats.Paths, path)
			}

			scannedPaths++
			onProgress(int64(scannedPaths * 100 / totalPaths))
		}

		if categoryStats.Size > 0 {
			result.Categories[category] = categoryStats
		}
	}

	return result, nil
}

func calculateDirSize(path string) (int64, int64, error) {
	var size int64
	var count int64
	semaphore := make(chan struct{}, 10)

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			size += info.Size()
			count++
		}
		return nil
	})

	_ = semaphore
	return size, count, err
}

func getCachePaths() map[string][]string {
	paths := make(map[string][]string)

	if runtime.GOOS == "windows" {
		tempDir := os.Getenv("TEMP")
		if tempDir == "" {
			tempDir = os.Getenv("TMP")
		}
		winDir := os.Getenv("WINDIR")
		appData := os.Getenv("APPDATA")

		paths["Temporary Files"] = []string{tempDir}
		if winDir != "" {
			paths["Windows Cache"] = []string{
				filepath.Join(winDir, "Temp"),
				filepath.Join(winDir, "SoftwareDistribution", "Download"),
			}
		}
		if appData != "" {
			paths["Application Cache"] = []string{
				filepath.Join(appData, "Local", "Temp"),
			}
		}
	} else if runtime.GOOS == "darwin" {
		home, _ := os.UserHomeDir()
		paths["User Cache"] = []string{
			filepath.Join(home, "Library", "Caches"),
		}
		paths["Trash"] = []string{
			filepath.Join(home, ".Trash"),
		}
		paths["System Temp"] = []string{
			"/var/tmp",
			"/tmp",
		}
	} else {
		home, _ := os.UserHomeDir()
		paths["User Cache"] = []string{
			filepath.Join(home, ".cache"),
		}
		paths["System Temp"] = []string{
			"/tmp",
			"/var/tmp",
		}
		if home != "" {
			paths["Downloads"] = []string{
				filepath.Join(home, "Downloads"),
			}
		}
	}

	return paths
}
