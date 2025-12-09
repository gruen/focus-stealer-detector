# Focus Stealer

A simple tool that monitors and logs which application has focus on macOS.

The Problem: Something I recently installed on MacOS kept stealing my windows focus while typing.

The Answer: This script.

**Spoiler:** It was Logitech. Get your shit together, fam. `/Applications/Utilities/LogiPluginService.app`

## Features

- Logs application focus changes with timestamps
- Shows the executable path of each focused application
- Configurable polling interval
- Lightweight and fast

## Installation

### Download Binary

Download the latest release from the [Releases](../../releases) page.

### Build from Source

Requires Go 1.21+ installed.

```bash
go build -o focus-stealer main.go
```

## Usage

```bash
# Run with default settings (50ms polling interval)
./focus-stealer

# Custom polling interval
./focus-stealer -interval 100ms
./focus-stealer -interval 1s
```

### Output

```
2025-12-09 14:23:45.123  Terminal  /System/Applications/Utilities/Terminal.app/
2025-12-09 14:23:47.456  Safari  /Applications/Safari.app/
2025-12-09 14:23:52.789  Code  /Applications/Visual Studio Code.app/
```

## Requirements

- macOS 10.15+ (uses AppleScript via `osascript`, included with macOS)

### Permissions

On first run, macOS will prompt for Accessibility permissions. Grant access to the terminal app you're using (Terminal, iTerm, etc.):

**System Settings → Privacy & Security → Accessibility** → Enable your terminal app

**Note:** Pre-built binaries have no external dependencies. Building from source requires only Go.

## Author

Michael E. Gruen ([michaelgruen.com](https://michaelgruen.com))

## License

MIT License - Copyright (c) 2025 Michael E. Gruen
