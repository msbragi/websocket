# TODO.md

## WebSocket File Transfer Project: Next Steps

This document tracks improvements, refactoring, and features for both the static server and websocket backend, following the new cross-platform, config-driven pattern.

---

### 1. **HTTPS & Certificates**
- [ ] Add HTTPS support to the websocket server (`ListenAndServeTLS`)
- [ ] Generate and use self-signed certificates for local development
- [ ] Update frontend to use `wss://` for secure websocket connections
- [ ] Document certificate setup for production

---

### 2. **Configurable Uploads Directory & Server Settings**
- [ ] Use a config file (JSON, TOML, or YAML) for all server settings (port, uploads directory, etc.)
- [ ] Create a separate config/setup application (GUI or CLI) for easy configuration
- [ ] Validate uploads directory exists and is writable on startup
- [ ] Allow changing config via the config app

---

### 3. **Installer & Autostart**
- [ ] Create an installer for Windows and Linux/macOS
- [ ] Installer should unpack websocket server and config app
- [ ] On Windows, create shortcut and add to autorun
- [ ] On Linux/macOS, create autostart entry (`.desktop` file)
- [ ] Document installation and autostart steps

---

### 4. **Frontend Enhancements**
- [ ] Add better error handling and user feedback in the UI
- [ ] Support drag-and-drop file uploads
- [ ] Show upload/download progress bars
- [ ] Add authentication or access control (optional)

---

### 5. **Deployment & Real-World Usage**
- [ ] Document how to run static server and websocket server on separate hosts
- [ ] Support remote static hosting (e.g., `https://example.com/mysite`)
- [ ] Ensure websocket server works with local connections (`wss://127.0.0.1/ws`)
- [ ] Add instructions for port forwarding and firewall configuration

---

### 6. **Codebase & Structure**
- [ ] Refactor code for clarity and maintainability
- [ ] Add unit and integration tests
- [ ] Document all public functions and APIs
- [ ] Clean up unused dependencies in `go.mod`

---

### 7. **General Improvements**
- [ ] Add logging and diagnostics for easier troubleshooting
- [ ] Provide sample config files and example icons
- [ ] Write a comprehensive README with setup, usage, and troubleshooting

---

### 8. **Future Improvements**
- [ ] Add SQLite database support for audit and configuration
- [ ] Add update checker to config app
- [ ] Show server status in config app
- [ ] Support backup/restore of config and uploads

---

**Note:**
- Tray icon functionality has been removed for cross-platform compatibility.
- All configuration is handled via a separate config/setup application and config file.
- Installer and autostart features will ensure smooth user experience on all platforms.

**Feel free to add, edit, or reorder items as the project evolves!**

# Create windows executable
$ GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o <filename>.exe