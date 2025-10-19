package system

import (
	"runtime"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsage(path string) (uint64, uint64, uint64, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return 0, 0, 0, err
	}

	return usage.Total, usage.Used, usage.Free, nil
}

func GetDiskUsagePercentage(path string) (float64, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return 0, err
	}

	return usage.UsedPercent, nil
}

func GetSystemDrive() string {
	if runtime.GOOS == "windows" {
		return "C:\\"
	}
	return "/"
}
