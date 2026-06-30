---
name: vikunja-cli
description: >
  Manage tasks, projects, teams, labels, and more by running the bundled `vikunja-cli`
  command-line tool against a Vikunja instance. Use whenever the user mentions Vikunja,
  tasks, todos, projects, labels, teams, or any project/task management activity. Covers
  creating/listing/updating/deleting tasks, projects, labels, teams, comments, assignees,
  attachments, notifications, filters, reactions, webhooks, views, buckets, shares,
  subscriptions, tokens, and user settings.
---

# Vikunja CLI

`vikunja-cli` is a CLI that wraps the Vikunja project-management API. All output is JSON.

## Setup — get the binary (first run only)

This skill ships its instructions but **not** the `vikunja-cli` binary (binaries are
large and platform-specific). Before the first call, make sure the binary exists at
`bin/vikunja-cli` inside this skill's directory. If it's missing, run the bundled
installer **once** — it downloads the right build for this machine from the GitHub
release and verifies its SHA256 checksum:

```sh
bash <skill-dir>/scripts/install-binary.sh
```

It writes `bin/vikunja-cli` (and prints the version), after which the binary persists —
you don't run this again. If the download is blocked, the fallbacks are
`go install github.com/jo-nike/vikunja-cli@latest` (Go 1.25+) or grabbing a binary from
<https://github.com/jo-nike/vikunja-cli/releases/latest> and dropping it at
`bin/vikunja-cli`.

## Running the binary

The binary lives at `bin/vikunja-cli` inside this skill's directory (the absolute base
directory is shown to you when the skill loads). Shell environment does **not** persist
between separate command calls, so use the full path on every invocation. If a
`vikunja-cli` is already on the user's `PATH`, that's the same tool and you can call it
bare — but the skill's own `bin/vikunja-cli` is the reliable default. If `bin/vikunja-cli`
is missing, do the **Setup** step above first. The examples below use a bare
`vikunja-cli` for readability.

Config: `~/.config/vikunja-cli/config.toml`

**Important:** API tokens (`tk_...`) only work for reads. For write operations (task updates, moves, bucket changes), configure username/password for JWT auth:

```toml
url = "https://vikunja.example.com"
username = "your-username"
password = "your-password"
```

The CLI auto-authenticates via JWT and caches it to `~/.cache/vikunja-cli/jwt-<hash>`. Falls back to `token` field if no credentials are set.

Env vars: `VIKUNJA_URL`, `VIKUNJA_TOKEN`, `VIKUNJA_USERNAME`, `VIKUNJA_PASSWORD`

## Command reference

Read [references/command-reference.md](references/command-reference.md) for the full command reference with all flags and API endpoints. Every command group is a top-level `##` heading, and each subcommand a `###` heading — grep for `## <group>` to jump to a section or `### <group> <verb>` for a specific command:

- `## tasks` — create, get, list, update, delete, bulk, move, assignees, attachments, comments, labels, relations
- `## projects` — create, get, list, update, delete, duplicate, backgrounds, buckets, shares, teams, users, views, webhooks
- `## teams` — create, get, list, update, delete, members
- `## labels` — create, get, list, update, delete
- `## filters` — saved filters
- `## notifications` — list, read, read-all
- `## reactions` — reactions on tasks and comments
- `## subscriptions` — subscribe/unsubscribe to projects and tasks
- `## migration` — import from Todoist, Trello, Microsoft To-Do, TickTick
- `## tokens` — API tokens
- `## user` — account, settings, TOTP, export, deletion, CalDAV tokens
- `## auth` — login, register
- `## system` — instance info

## Common workflows

### List tasks
```
vikunja-cli tasks list
vikunja-cli tasks list --project-id 1
vikunja-cli tasks list --filter "done = false" --sort due_date --order-by asc
vikunja-cli tasks list --search "meeting"
vikunja-cli tasks list --project-id 1 --bucket backlog
vikunja-cli tasks list --project-id 1 --bucket-id 42 --view-id 7
```

### Create a task
```
vikunja-cli tasks create --project-id 1 --title "Buy groceries" --priority 3 --due-date "2025-03-01T17:00:00Z"
```

### Complete a task
```
vikunja-cli tasks update --id 42 --done
```

### Move a task to a bucket
```
vikunja-cli tasks move --id 42 --bucket "In Progress" --project-id 1
vikunja-cli tasks move --id 42 --bucket-id 3
```

### Create a task in a specific bucket
```
vikunja-cli tasks create --project-id 1 --title "Review PR" --bucket "In Progress"
vikunja-cli tasks create --project-id 1 --title "Review PR" --bucket "Todo" --view-id 2
```

### Bulk operations
```
vikunja-cli tasks bulk --task-ids "1,2,3" --done
vikunja-cli tasks bulk --task-ids "1,2,3" --bucket "Done" --project-id 1
vikunja-cli tasks labels bulk --task-id 5 --label-ids "1,2,3"
vikunja-cli tasks assignees bulk --task-id 5 --user-ids "1,2"
```

### Project management
```
vikunja-cli projects list
vikunja-cli projects create --title "Q1 Planning" --description "Quarterly goals"
vikunja-cli projects get --id 1
```

## Text formatting (IMPORTANT!)

**Vikunja requires HTML line breaks for proper formatting:**

❌ **Wrong** (plain newlines get mangled):
```bash
vikunja-cli tasks create --project-id 1 --title "Task" --description "Line 1
Line 2
Line 3"
```

✅ **Right** (use `<br>` for line breaks):
```bash
vikunja-cli tasks create --project-id 1 --title "Task" --description "Line 1<br>Line 2<br>Line 3"
```

✅ **Also right** (use `<br><br>` for paragraph breaks):
```bash
vikunja-cli tasks comments create --task-id 42 --comment "Summary line<br><br>Details:<br>- Point 1<br>- Point 2"
```

**Apply to all text fields:** `--description`, `--comment`, task updates, etc.

**Markdown formatting also works:**
- **Bold:** `**text**`
- *Italic:* `*text*`
- Lists: `- item` or `1. item`
- But still use `<br>` for explicit line breaks within paragraphs

## Key conventions

- IDs are integers passed via `--id`, `--task-id`, `--project-id`, etc.
- Dates use RFC3339 format: `2025-03-01T17:00:00Z`
- Priority range: 0 (none) to 5 (highest)
- Percent done: 0.0 to 1.0
- Rights: 0=read, 1=read&write, 2=admin
- Boolean flags like `--done`, `--is-favorite` are toggles
- Pagination: `--page` and `--per-page` on list commands
- `--bucket` accepts a bucket name (resolved to ID); use `--view-id` to disambiguate if multiple views have same-named buckets
- Parse JSON output with `jq` for extracting specific fields
