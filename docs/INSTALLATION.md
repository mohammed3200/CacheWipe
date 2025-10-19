# Installation Guide

## System Requirements

- OS: Windows 7+, macOS 10.13+, Linux
- RAM: 256MB minimum
- Disk: 50MB free space
- Go 1.21+ (for building from source)

## Installation Methods

### 1. Pre-built Binaries

Download from [Releases](https://github.com/mohammed3200/cache-wipe/releases)

**Linux/macOS:**
```bash
chmod +x cache-wipe-linux-amd64
./cache-wipe-linux-amd64
```

**Windows:**
- Double-click `cache-wipe-windows-amd64.exe`

### 2. From Source
```bash
git clone https://github.com/mohammed3200/cache-wipe.git
cd cache-wipe
go build
./cache-wipe
```

### 3. Using Install Script

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/mohammed3200/cache-wipe/main/install.sh | bash
```

### 4. Docker
```bash
docker pull mohammed3200/cache-wipe:latest
docker run -it mohammed3200/cache-wipe:latest
```

## Uninstallation

**Linux/macOS:**
```bash
sudo rm /usr/local/bin/cache-wipe
```

**Windows:**
- Use Add/Remove Programs

## Troubleshooting

### Application crashes on startup
- Ensure all dependencies are installed
- Check system logs
- Rebuild from source

### Permission denied
- Run with sudo/administrator privileges
- Check file permissions

### Missing library errors
- Install required system libraries
- Update Go to latest version