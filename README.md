# ghprofile

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/go-1.25+-00ADD8.svg)](https://golang.org/)

A beautiful, customizable GitHub profile viewer for your terminal.

![ghprofile demo](https://github.com/user-attachments/assets/6e0aff20-7d3d-44af-8703-97c202a52199)


## Features

- **Profile & repos**: Fetch user profiles, top repos, language stats
- **Contribution graph**: Shows activity (requires token)
- **Compare users**: Side-by-side profile comparison
- **Flexible output**: JSON, text, TUI, or plain formats
- **Smart caching**: Offline mode + rate-limit awareness
- **Customizable**: Icons, borders, sizes, sorting, filtering
- **Config file**: Persistent token storage

## Quick Start

```sh
go build -o ghprofile ./cmd/main.go
./ghprofile --user dayvster
```

With token (higher rate limit + contribution graph):
```sh
./ghprofile --user dayvster --token ghp_xxxxxxxxxxxx
```

## Usage

### Common Options

```sh
./ghprofile --user <username> [flags]
```

Examples:
```sh
./ghprofile --user dayvster --json              # JSON output
./ghprofile --user user1,user2                  # Compare multiple
./ghprofile --demo                              # Demo mode (no network)
./ghprofile --user dayvster --sort forks        # Sort by forks
```

### All Flags

| Flag | Default | Description |
|------|---------|-------------|
| Flag | Default | Description |
|------|---------|-------------|
| `-u`, `--user` | — | GitHub username(s) (comma-separated for multiple) |
| `-n` | `5` | Top repos to show |
| `--token` | — | GitHub personal access token (or `GITHUB_TOKEN` env) |
| `--sort` | `stars` | Sort by: `stars`, `forks`, `name`, `updated` |
| `--json` | `false` | Output as JSON |
| `--tui` | `true` | Interactive TUI mode |
| `--demo` | `false` | Demo mode (no network) |
| `--no-icons` | `false` | Disable icons |
| `--no-border` | `false` | Remove card border |
| `--no-style` | `false` | Remove all styling |
| `--size` | `medium` | Output size: `small`, `medium`, `large`, `full` |
| `--since` | — | Only repos updated after date (`YYYY-MM-DD`) |
| `--verbose`, `-v` | `false` | Print API details to stderr |
| `--quiet`, `-q` | `false` | Suppress stderr warnings |
| `--help`, `-h` | — | Show help |

## Configuration

Token resolution order:
1. `--token` flag
2. `GITHUB_TOKEN` environment variable
3. Config file: `~/.config/ghprofile/config.json`

Save your token:
```sh
mkdir -p ~/.config/ghprofile
echo '{"token": "ghp_xxxxxxxxxxxx"}' > ~/.config/ghprofile/config.json
```

## License

See [LICENSE](LICENSE).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).
