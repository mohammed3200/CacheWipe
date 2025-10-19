# Cache wipe Desktop Application

A lightweight, cross-platform desktop application to help you clean cache files and optimize your computer's performance. Built with Go and Fyne framework.

## Features

- **Multi-Platform Support**: Runs on Windows, macOS, and Linux
- **System Cache Detection**: Automatically finds cache locations
- **Real-Time Scanning**: Quick and efficient cache analysis
- **One-Click Cleanup**: Remove cache with a single button
- **Detailed Statistics**: View cache size and file count
- **Safe Deletion**: Backup logs of all deletions
- **Progress Tracking**: Visual feedback during operations
- **Cross-Platform GUI**: Native-looking interface on all platforms

## Supported Cache Types

- Windows: Temp files, Windows cache, application data
- macOS: User library cache, trash, system temp files
- Linux: .cache directory, system temp files

## Requirements

- Go 1.21 or higher
- 50MB free disk space
- Administrator/sudo privileges for system cache cleanup

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/yourusername/cache-wipe.git
cd cache-wipe
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o cache-wipe
```

4. Run the application:
```bash
./cache-wipe
```

### Pre-built Binaries

Download pre-built executables from the [Releases](https://github.com/yourusername/cache-wipe/releases) page.

## Usage

### GUI Application

1. Launch the application
2. Click "Scan Now" to analyze your system cache
3. Review the detected cache locations
4. Click "Clean Cache" to remove selected cache files
5. Check the summary report

### Command Line (Optional)

```bash
# Scan cache
./cache-wipe scan

# Cleanup cache
./cache-wipe cleanup

# Show help
./cache-wipe help
```

## Project Structure

```
cache-wipe/
├── main.go                     # Entry point
├── go.mod                      # Module definition
├── go.sum                      # Dependencies
├── README.md                   # This file
├── LICENSE                     # MIT License
├── CONTRIBUTING.md             # Contribution guidelines
├── .gitignore                  # Git ignore file
├── internal/
│   ├── cache/
│   │   ├── wipe.go         # Cleanup logic
│   │   └── scanner.go         # Cache scanning
│   ├── ui/
│   │   ├── window.go          # Main UI window
│   │   ├── widgets.go         # Custom widgets
│   │   └── styles.go          # UI styling
│   └── system/
│       ├── disk.go            # Disk utilities
│       └── platform.go        # Platform detection
├── pkg/
│   └── models/
│       └── stats.go           # Data models
├── resources/
│   └── icon.png               # Application icon
└── build/                     # Build outputs
```

## Architecture

### Core Modules

**Cache Module** - Handles scanning and cleaning operations
- Scanner: Identifies cache locations and calculates sizes
- wipe: Safely removes cache files

**UI Module** - Manages user interface
- Window: Main application window setup
- Widgets: Custom UI components
- Styles: Theme and styling

**System Module** - Platform-specific utilities
- Platform detection (Windows/macOS/Linux)
- Disk usage information
- Cache path resolution

**Models** - Data structures for statistics and configuration

## Development

### Setup Development Environment

```bash
# Clone repository
git clone https://github.com/mohammed3200/cache-wipe.git
cd cache-wipe

# Install dependencies
go mod download

# Run with live reload (optional, using air)
air
```

### Building

```bash
# Standard build
go build -o cache-wipe

# Optimized build (smaller size)
go build -ldflags="-s -w" -o cache-wipe

# Cross-compile for Windows from Linux/macOS
GOOS=windows GOARCH=amd64 go build -o cache-wipe.exe

# Cross-compile for macOS from Linux/Windows
GOOS=darwin GOARCH=amd64 go build -o cache-wipe-mac

# Cross-compile for Linux from Windows/macOS
GOOS=linux GOARCH=amd64 go build -o cache-wipe-linux
```

### Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific test
go test -run TestScannerScan ./internal/cache

# Run with coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Code Quality

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Get code report
go vet ./...
```

## Security Considerations

- **No Root Required**: Most operations work with user permissions
- **Confirmation Dialogs**: Always confirm before deleting files
- **Deletion Logs**: Maintain audit trail of removed files
- **Path Validation**: Prevent directory traversal attacks
- **Backup**: Option to backup cache before deletion

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.

## Troubleshooting

### Application won't start
- Ensure Go 1.21+ is installed
- Try rebuilding: `go clean && go build`
- Check system logs for errors

### Permission denied errors
- Run with administrator/sudo privileges
- Some system cache requires elevated permissions

### No cache detected
- System cache may already be clean
- Check if cache paths exist on your system
- Review application logs

## Performance Tips

- Run cleanup regularly (weekly recommended)
- Close resource-heavy applications before cleanup
- Schedule automatic cleanup during off-hours
- Monitor disk usage improvements

## Roadmap

### v1.0 (Current)
- Basic cache scanning and cleanup
- Multi-platform support
- Simple GUI

### v1.1 (Planned)
- Selective cache cleanup
- Cleanup scheduling
- Cleanup history

### v1.2 (Future)
- Browser cache detection
- Duplicate file finder
- System optimization tips

## Support

- Report bugs via [GitHub Issues](https://github.com/yourusername/cache-wipe/issues)
- Start a [GitHub Discussion](https://github.com/yourusername/cache-wipe/discussions)
- Check [Documentation](https://github.com/yourusername/cache-wipe/wiki)

## Credits

Built with:
- [Go](https://golang.org/)
- [Fyne](https://fyne.io/) - Cross-platform GUI toolkit

## Changelog

### v1.0 (2024)
- Initial release
- Windows, macOS, Linux support
- Basic cache cleaning functionality
- GUI application

---

**Star this project if you find it useful!** ⭐

---