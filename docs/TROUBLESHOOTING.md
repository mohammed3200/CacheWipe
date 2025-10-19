# Troubleshooting Guide

## Common Issues

### Application Won't Start

**Symptoms:** Crash on launch or no window appears

**Solutions:**
1. Check system requirements (Go 1.21+)
2. Reinstall application
3. Check system logs
4. Try running with debug output

**Debug:**
```bash
./cache-wipe -debug
```

### Permission Denied Errors

**Symptoms:** "Access denied" when cleaning system cache

**Solutions:**
1. Run with administrator/sudo:
   - Linux/macOS: `sudo ./cache-wipe`
   - Windows: Right-click → "Run as administrator"
2. Check file permissions
3. Disable system cache cleanup temporarily

### Scan Takes Too Long

**Symptoms:** Scanning appears stuck

**Solutions:**
1. Cancel and try again
2. Reduce cache paths
3. Check disk health
4. Close other applications

### No Cache Found

**Symptoms:** "Total Cache: 0 B" after scan

**Solutions:**
1. System cache may be clean
2. Cache paths might not exist
3. Insufficient permissions
4. Try running as administrator

### High Memory Usage

**Symptoms:** Application uses excessive RAM

**Solutions:**
1. Close other applications
2. Reduce cache path count
3. Restart computer
4. Update to latest version

## Log Analysis

Check logs for detailed error information:
- Linux/macOS: `~/.cache/cache-wipe/`
- Windows: `%APPDATA%\cache-wipe\`

## Getting Help

1. Check documentation at: https://github.com/mohammed3200/cache-wipe/wiki
2. Search existing issues: https://github.com/mohammed3200/cache-wipe/issues
3. Create new issue with:
   - OS and version
   - Error message
   - Steps to reproduce
   - Attached logs