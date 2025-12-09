# focus-stealer

A simple CLI tool that monitors and logs which application has focus on macOS.

## Features

- Logs application focus changes with timestamps
- Shows the executable path of each focused application
- Configurable polling interval
- Lightweight and fast

## Installation

### Download Binary

Download the latest release from the [Releases](../../releases) page.

### Build from Source

```bash
go build -o focus-stealer app.go
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

- macOS (uses AppleScript via `osascript`)
- Accessibility permissions for Terminal/iTerm (grant when prompted)

## Author

Michael E. Gruen ([michaelgruen.com](https://michaelgruen.com))

## License

MIT License - Copyright (c) 2025 Michael E. Gruen
