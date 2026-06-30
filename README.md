# vikunja-cli

A command-line interface for [Vikunja](https://vikunja.io), the open-source to-do
and project-management app. Every command prints JSON, so it composes cleanly with
`jq`, scripts, and other tooling.

## Installation

### Prebuilt release (recommended)

Download the archive for your platform from the
[latest release](https://github.com/jo-nike/vikunja-cli/releases/latest), verify
its checksum against `checksums.txt`, and extract the binary onto your `PATH`:

```bash
# macOS (Apple Silicon) example
curl -fsSLO https://github.com/jo-nike/vikunja-cli/releases/latest/download/vikunja-cli-darwin-arm64.tar.gz
curl -fsSLO https://github.com/jo-nike/vikunja-cli/releases/latest/download/checksums.txt
shasum -a 256 -c checksums.txt --ignore-missing
tar -xzf vikunja-cli-darwin-arm64.tar.gz
sudo mv vikunja-cli /usr/local/bin/
```

Available archives: `darwin-{amd64,arm64}`, `linux-{amd64,arm64}`, `windows-amd64`
(`.zip`). On macOS, Gatekeeper may quarantine the unsigned binary on first run —
clear it with `xattr -d com.apple.quarantine /usr/local/bin/vikunja-cli`. (Binaries
fetched via `curl` are generally not quarantined.) On Windows, use "More info → Run
anyway" or unblock the `.zip` in its file properties.

### With Go

```bash
go install github.com/jo-nike/vikunja-cli@latest
```

### From source

```bash
git clone https://github.com/jo-nike/vikunja-cli.git
cd vikunja-cli
make build      # produces ./vikunja-cli with version metadata baked in
```

## Configuration

Create `~/.config/vikunja-cli/config.toml`:

```toml
url = "https://vikunja.example.com"
token = "your-api-token"
```

### JWT authentication (recommended for writes)

Vikunja API tokens (`tk_...`) work for reads but silently ignore many write
operations (task updates, moves, bucket changes). For full read/write support,
configure a username and password instead:

```toml
url = "https://vikunja.example.com"
username = "your-username"
password = "your-password"
```

With credentials set, the CLI logs in automatically to obtain a JWT and caches it
under `~/.cache/vikunja-cli/` to avoid re-authenticating on every call. The `token`
field is used as a fallback when no credentials are configured.

### Environment variables

Environment variables override config-file values:

| Variable | Description |
|---|---|
| `VIKUNJA_URL` | Base URL of the Vikunja instance |
| `VIKUNJA_TOKEN` | API token or JWT for authentication |
| `VIKUNJA_USERNAME` | Username for JWT authentication |
| `VIKUNJA_PASSWORD` | Password for JWT authentication |

## Quickstart

```bash
# Authenticate (prints a JWT; usually you just set credentials in config instead)
vikunja-cli auth login --username user --password pass

# Projects
vikunja-cli projects list
vikunja-cli projects create --title "Q1 Planning" --description "Quarterly goals"

# Tasks
vikunja-cli tasks create --project-id 1 --title "My task" --priority 3
vikunja-cli tasks list --filter "done = false" --sort due_date --order-by asc
vikunja-cli tasks update --id 42 --done

# Labels
vikunja-cli labels create --title "urgent" --hex-color "#ff0000"

# Instance info
vikunja-cli system info
```

All output is JSON — pipe it through `jq` to extract fields:

```bash
vikunja-cli tasks list --project-id 1 | jq '.[] | {id, title, done}'
```

> **Formatting note:** Vikunja stores rich text as HTML. Use `<br>` for line
> breaks (and `<br><br>` for paragraphs) in `--description`/`--comment`; plain
> newlines get mangled. Markdown (`**bold**`, `- list`) works too.

### Available commands

| Command | Description |
|---|---|
| `auth` | Login and register |
| `projects` | Projects, views, buckets, shares, webhooks, backgrounds, and team/user members |
| `tasks` | Tasks, assignees, attachments, comments, labels, and relations |
| `labels` | Labels |
| `teams` | Teams and team members |
| `tokens` | API tokens |
| `filters` | Saved filters |
| `notifications` | List and mark notifications as read |
| `reactions` | Reactions on tasks and comments |
| `subscriptions` | Subscribe/unsubscribe to projects and tasks |
| `migration` | Import from Todoist, Trello, Microsoft To-Do, TickTick |
| `user` | Account, settings, TOTP, CalDAV tokens, export, and deletion |
| `system` | Vikunja instance info |
| `version` | Print CLI version |

Run `vikunja-cli <command> --help` for the flags of any command, and see
[DOCS.md](DOCS.md) for the full command reference with every flag and API endpoint.

### Shell completion

```bash
source <(vikunja-cli completion bash)   # Bash
source <(vikunja-cli completion zsh)    # Zsh
vikunja-cli completion fish | source    # Fish
```

## Use it from Claude Code

This repo ships a Claude Code skill so an agent can drive Vikunja for you. Install
it with the [`skills`](https://github.com/anthropics/skills) CLI:

```bash
npx skills@latest add jo-nike/vikunja-cli -a claude-code --copy
```

On first use the skill downloads the matching binary into its own `bin/` (see
`.claude/skills/vikunja-cli/scripts/install-binary.sh`); configure auth as above.

## Acknowledgments

This project talks to the [Vikunja](https://vikunja.io) API. Vikunja is an
open-source project-management application licensed under the
[GNU AGPLv3](https://github.com/go-vikunja/vikunja/blob/main/LICENSE).

## License

Licensed under the MIT License. See [LICENSE](LICENSE) for details.
