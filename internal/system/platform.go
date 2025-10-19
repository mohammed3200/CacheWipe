package system

import (
	"os"
	"path/filepath"
	"runtime"
)

type Platform string

const (
	Windows Platform = "windows"
	Darwin  Platform = "darwin"
	Linux   Platform = "linux"
)

func GetCurrentPlatform() Platform {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "darwin":
		return Darwin
	case "linux":
		return Linux
	default:
		return Linux
	}
}

func GetHomeDir() (string, error) {
	return os.UserHomeDir()
}

func GetCacheDir() string {
	home, _ := GetHomeDir()
	platform := GetCurrentPlatform()

	switch platform {
	case Windows:
		return filepath.Join(os.Getenv("APPDATA"), "Local", "Temp")
	case Darwin:
		return filepath.Join(home, "Library", "Caches")
	default:
		return filepath.Join(home, ".cache")
	}
}
