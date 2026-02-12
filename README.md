# vikunja-cli

A command-line interface for [Vikunja](https://vikunja.io), the open-source to-do and project management application. All output is JSON, making it easy to integrate with scripts and other tools.

## Installation

```bash
go install github.com/jo-nike/vikunja-cli@latest
```

Or build from source:

```bash
git clone https://github.com/jo-nike/vikunja-cli.git
cd vikunja-cli
make build
```

## Configuration

Create a config file at `~/.config/vikunja-cli/config.toml`:

```toml
url = "https://vikunja.example.com"
token = "your-api-token"
```

Or use environment variables:

| Variable | Description |
|---|---|
| `VIKUNJA_URL` | Base URL of the Vikunja instance |
| `VIKUNJA_TOKEN` | API token or JWT for authentication |

## Usage

```bash
# Login and get a JWT token
vikunja auth login --username user --password pass

# List all projects
vikunja projects list

# Create a task
vikunja tasks create --project-id 1 --title "My task"

# List tasks with filtering
vikunja tasks list --filter "done = false"

# Manage labels
vikunja labels create --title "urgent" --hex-color "#ff0000"

# Get system info
vikunja system info
```

### Available Commands

| Command | Description |
|---|---|
| `auth` | Login and register |
| `projects` | Manage projects, views, buckets, shares, webhooks, backgrounds, and team/user members |
| `tasks` | Manage tasks, assignees, attachments, comments, labels, and relations |
| `labels` | Manage labels |
| `teams` | Manage teams and team members |
| `tokens` | Manage API tokens |
| `filters` | Manage saved filters |
| `notifications` | List and mark notifications as read |
| `reactions` | Manage reactions on tasks and comments |
| `subscriptions` | Subscribe/unsubscribe to projects and tasks |
| `migration` | Import data from Todoist, Trello, Microsoft To-Do, TickTick |
| `user` | Manage account, settings, TOTP, CalDAV tokens, export, and deletion |
| `system` | Get Vikunja instance info |
| `version` | Print CLI version |

Run `vikunja [command] --help` for detailed usage of any command.

### Shell Completion

```bash
# Bash
source <(vikunja completion bash)

# Zsh
source <(vikunja completion zsh)

# Fish
vikunja completion fish | source
```

See [DOCS.md](DOCS.md) for the full command reference.

## Acknowledgments

This project interacts with the [Vikunja](https://vikunja.io) API. Vikunja is an open-source project management application licensed under the [GNU AGPLv3](https://github.com/go-vikunja/vikunja/blob/main/LICENSE).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
