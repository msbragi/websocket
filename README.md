# Websocket Server Project

This repository contains a Go-based WebSocket server with a browser-accessible configuration and file management interface.

## Features

- **WebSocket Server:** Handles real-time communication via WebSockets.
- **Configurable:** Easily update server settings (port, upload directory, update URL) via a web interface.
- **File Uploads:** Supports file uploads and downloads.
- **Directory Browser:** Browse, create, and delete folders from the web UI.
- **Static Web UI:** User-friendly HTML/CSS interface for configuration and file management.

## Project Structure

- [`Server/`](websocket/Server/): Main Go server and static web assets.
  - `main.go`: Server entry point.
  - `static/`: Contains `index.html`, `config.html`, and `style.css`.
- [`ws-main/`](websocket/ws-main/): WebSocket handlers and utilities.
  - `main.go`: WebSocket server entry point.
  - `browser_handler.go`: Directory browsing logic.
  - `config_handler.go`: Configuration management.
  - `file_handler.go`: File upload/download logic.
  - `utils.go`: Shared utilities.
  - `config/`: Stores `config.json` for server settings.
  - `uploads/`: Default upload directory.
  - `assets/`: Icons and other assets.

## Getting Started

### Prerequisites

- Go 1.22+
- [Gorilla WebSocket](https://github.com/gorilla/websocket) package

### Build & Run

```sh
cd websocket/Server
go build -o websocket-server main.go
./websocket-server
```

### Access the Web UI

Open [http://localhost:8088/static/index.html](http://localhost:8088/static/index.html) in your browser.

## Configuration

Edit [`config/config.json`](websocket/ws-main/config/config.json) or use the web UI at `/static/config.html` to update:

- `server_port`: Port for the WebSocket server.
- `upload_dir`: Directory for file uploads.
- `update_url`: URL for updates.

## Development

This project is ready for development in a VS Code dev container (see [`.devcontainer/devcontainer.json`](websocket/.devcontainer/devcontainer.json)).

## License

MIT

---

**TODO:** See [`TODO.md`](websocket/TODO.md) for planned features