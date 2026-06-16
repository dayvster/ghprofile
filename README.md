<div align="center">

# ghprofile

A beautiful, customizable GitHub profile viewer for your terminal.

---

<p>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
  <img src="https://img.shields.io/badge/go-1.25+-00ADD8.svg" alt="Go Version">
  <img src="https://img.shields.io/github/stars/dayvster/ghprofile?style=social" alt="GitHub stars">
  <img src="https://img.shields.io/github/downloads/dayvster/ghprofile/total.svg" alt="Downloads">
  <img src="https://img.shields.io/github/issues/dayvster/ghprofile.svg" alt="Issues">
  <img src="https://img.shields.io/github/forks/dayvster/ghprofile.svg" alt="Forks">
  <img src="https://img.shields.io/github/repo-size/dayvster/ghprofile.svg" alt="Repo size">
  <img src="https://img.shields.io/github/last-commit/dayvster/ghprofile.svg" alt="Last commit">
</p>

</div>

---

<img width="1426" height="1083" alt="image" src="https://github.com/user-attachments/assets/6e0aff20-7d3d-44af-8703-97c202a52199" />


## Features

- Fetch and display GitHub user profiles and top repositories
- Language stats, repo stars, forks, and more
- Contribution graph (requires GitHub token)
- Compare multiple users at once
- Icon-rich output with Nerd Font support
- JSON output for piping
- Caching and demo mode for offline/limited API use
- Config file for persistent token storage
- TUI mode with Bubble Tea

---

## Installation

```sh
go build -o ghprofile ./cmd/main.go
```

---

## Usage

```sh
./ghprofile --user <username> [flags]
```

### Basic

```sh
# Single user
./ghprofile --user dayvster

# With token (higher rate limit + contribution graph)
./ghprofile --user dayvster --token ghp_xxxxxxxxxxxx

# Compare multiple users
./ghprofile --user dayvster,torvalds

# JSON output
./ghprofile --user dayvster --json

# Plain text (no TUI, no icons, no borders)
./ghprofile --user dayvster --tui=false --no-icons --no-border

# Demo data (no network)
./ghprofile --demo
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-u`, `--user` | — | GitHub username(s) to fetch (comma-separated for multiple) |
| `-n` | `5` | How many top repos to show |
| `--token` | — | GitHub personal access token (or `GITHUB_TOKEN` env) |
| `--config` | `~/.config/ghprofile/config.json` | Config file path |
| `--json` | `false` | Output as JSON |
| `--format` | `text` | Output format: `text`, `json` |
| `--no-icons` | `false` | Disable icons in the output |
| `--no-border` | `false` | Remove card border from output |
| `--no-style` | `false` | Remove all styles from output |
| `--no-demo` | `false` | Exit on fetch error instead of falling back to demo data |
| `--demo` | `false` | Force demo data (skip network and cache) |
| `--size` | `medium` | Output size: `small`, `medium`, `large`, `full` |
| `--tui` | `true` | Use interactive TUI mode |
| `-h`, `--help` | — | Show help message |

### Configuration

Token resolution priority:
1. `--token` flag
2. `GITHUB_TOKEN` environment variable
3. Config file (`~/.config/ghprofile/config.json` or `$XDG_CONFIG_HOME/ghprofile/config.json`)

```sh
mkdir -p ~/.config/ghprofile
echo '{"token": "ghp_xxxxxxxxxxxx"}' > ~/.config/ghprofile/config.json
```

### Examples

```sh
# Contribution graph (requires token)
GITHUB_TOKEN=ghp_xxxxxxxxxxxx ./ghprofile --user dayvster

# JSON output piped to jq
./ghprofile --user dayvster --json | jq '.profile.followers_amount'

# Multiple users compared
./ghprofile --user dayvster,torvalds --no-icons --size small

# Custom config file
./ghprofile --user dayvster --config /path/to/config.json
```

---

## Screenshots

| Default | No Border | No Icons |
|---------|-----------|----------|
| ![ghprofile demo](screenshots/basic.png) | ![No border](screenshots/no-border.png) | ![No icons](screenshots/no-icons.png) |

---

## License

See [LICENSE](LICENSE).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).
