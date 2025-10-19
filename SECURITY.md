# Security Policy

## Reporting Security Issues

Please do not create GitHub issues for security vulnerabilities. Instead, email security@example.com with:

- Description of vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (optional)

## Security Best Practices

### For Users
1. Download from official sources only
2. Verify checksums
3. Run with minimal required permissions
4. Keep application updated
5. Enable cleanup logging
6. Review deletion logs regularly

### For Developers
1. Always validate file paths
2. Prevent directory traversal
3. Require confirmation for destructive operations
4. Maintain audit logs
5. Use secure coding practices
6. Regular security audits
7. Dependency scanning

## Known Limitations

- System cache cleanup requires elevated privileges
- Some cache locations are platform-specific
- Browser caches may require browsers to be closed
- Cloud-synced files should not be cleaned

## Security Updates

Check for updates regularly:
- GitHub Releases
- Official website
- Security advisories