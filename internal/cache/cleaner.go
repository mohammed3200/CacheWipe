package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type Cleaner struct {
	stopped int32
	mu      sync.Mutex
}

type CleanupProgress struct {
	Current      int64
	Total        int64
	DeletedSize  int64
	DeletedCount int64
	CurrentFile  string
	Status       string
}

type CleanupResult struct {
	DeletedSize  int64
	DeletedCount int64
	ErrorCount   int64
	Errors       []error
}

func NewCleaner() *Cleaner {
	return &Cleaner{
		stopped: 0,
	}
}

func (c *Cleaner) Stop() {
	atomic.StoreInt32(&c.stopped, 1)
}

func (c *Cleaner) IsStopped() bool {
	return atomic.LoadInt32(&c.stopped) == 1
}

func (c *Cleaner) Cleanup(items []CacheItem, onProgress func(CleanupProgress)) *CleanupResult {
	result := &CleanupResult{
		Errors: []error{},
	}

	atomic.StoreInt32(&c.stopped, 0)
	total := int64(len(items))

	for i, item := range items {
		if c.IsStopped() {
			break
		}

		if !item.Selected {
			continue
		}

		deletedSize, deletedCount, err := c.cleanPath(item.Path)
		if err != nil {
			result.Errors = append(result.Errors, err)
			result.ErrorCount++
		}

		result.DeletedSize += deletedSize
		result.DeletedCount += deletedCount

		onProgress(CleanupProgress{
			Current:      int64(i + 1),
			Total:        total,
			DeletedSize:  result.DeletedSize,
			DeletedCount: result.DeletedCount,
			CurrentFile:  item.Path,
			Status:       fmt.Sprintf("Cleaning %s", item.Category),
		})
	}

	return result
}

func (c *Cleaner) CleanupSpecific(paths []string, onProgress func(CleanupProgress)) *CleanupResult {
	result := &CleanupResult{
		Errors: []error{},
	}

	atomic.StoreInt32(&c.stopped, 0)
	total := int64(len(paths))

	for i, path := range paths {
		if c.IsStopped() {
			break
		}

		deletedSize, deletedCount, err := c.cleanPath(path)
		if err != nil {
			result.Errors = append(result.Errors, err)
			result.ErrorCount++
		}

		result.DeletedSize += deletedSize
		result.DeletedCount += deletedCount

		onProgress(CleanupProgress{
			Current:      int64(i + 1),
			Total:        total,
			DeletedSize:  result.DeletedSize,
			DeletedCount: result.DeletedCount,
			CurrentFile:  path,
		})
	}

	return result
}

func (c *Cleaner) cleanPath(path string) (int64, int64, error) {
	var deletedSize int64
	var deletedCount int64

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			deletedSize += info.Size()
			deletedCount++
			if err := os.Remove(filePath); err != nil {
				return nil
			}
		}
		return nil
	})

	if err == nil && deletedCount == 0 {
		if err := os.RemoveAll(path); err != nil {
			return 0, 0, err
		}
	}

	return deletedSize, deletedCount, nil
}

func (c *Cleaner) DryRun(items []CacheItem) (int64, int64) {
	totalSize := int64(0)
	totalCount := int64(0)

	for _, item := range items {
		if item.Selected {
			totalSize += item.Size
			totalCount += item.FileCount
		}
	}

	return totalSize, totalCount
}
