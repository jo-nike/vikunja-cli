# vikunja-cli Command Reference

A command-line interface for interacting with the Vikunja project management API. All output is JSON.

## Configuration

The CLI reads configuration from a TOML file and/or environment variables.

**Config file:** `~/.config/vikunja-cli/config.toml`

```toml
url = "https://vikunja.example.com"
token = "your-api-token"
```

**Environment variables** (override config file values):

| Variable | Description |
|---|---|
| `VIKUNJA_URL` | Base URL of the Vikunja instance |
| `VIKUNJA_TOKEN` | API token or JWT for authentication |

All API requests are made to `{VIKUNJA_URL}/api/v1/...`.

---

## auth

Authentication commands.

### auth login

Login and get a JWT token.

```
vikunja auth login [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--username` | string | Username or email (required) |
| `--password` | string | Password (required) |
| `--totp-passcode` | string | TOTP passcode |
| `--long-token` | bool | Request a long-lived token |

**API endpoint:** `POST /api/v1/login`

### auth register

Register a new user account.

```
vikunja auth register [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--username` | string | Username (required) |
| `--email` | string | Email (required) |
| `--password` | string | Password (required) |

**API endpoint:** `POST /api/v1/register`

---

## filters

Manage saved filters.

### filters create

Create a saved filter.

```
vikunja filters create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--title` | string | Filter title (required) |
| `--description` | string | Filter description |
| `--filters` | string | Filter definition as JSON |
| `--is-favorite` | bool | Mark as favorite |

**API endpoint:** `PUT /api/v1/filters`

### filters get

Get a saved filter by ID.

```
vikunja filters get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Filter ID (required) |

**API endpoint:** `GET /api/v1/filters/{id}`

### filters update

Update a saved filter.

```
vikunja filters update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Filter ID (required) |
| `--title` | string | Filter title |
| `--description` | string | Filter description |
| `--filters` | string | Filter definition as JSON |
| `--is-favorite` | bool | Mark as favorite |

**API endpoint:** `POST /api/v1/filters/{id}`

### filters delete

Delete a saved filter.

```
vikunja filters delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Filter ID (required) |

**API endpoint:** `DELETE /api/v1/filters/{id}`

---

## labels

Manage labels.

### labels create

Create a label.

```
vikunja labels create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--title` | string | Label title (required) |
| `--description` | string | Label description |
| `--hex-color` | string | Hex color (e.g. #ff0000) |

**API endpoint:** `PUT /api/v1/labels`

### labels get

Get a label by ID.

```
vikunja labels get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Label ID (required) |

**API endpoint:** `GET /api/v1/labels/{id}`

### labels list

List all labels.

```
vikunja labels list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--page` | int | Page number |
| `--per-page` | int | Items per page |
| `--search` | string | Search labels |

**API endpoint:** `GET /api/v1/labels`

### labels update

Update a label.

```
vikunja labels update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Label ID (required) |
| `--title` | string | Label title |
| `--description` | string | Label description |
| `--hex-color` | string | Hex color |

**API endpoint:** `POST /api/v1/labels/{id}`

### labels delete

Delete a label.

```
vikunja labels delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Label ID (required) |

**API endpoint:** `DELETE /api/v1/labels/{id}`

---

## migration

Import data from other services. Supported services: `todoist`, `trello`, `microsoft-todo`, `ticktick`, `vikunja-file`.

### migration auth

Get the auth URL for a migration service.

```
vikunja migration auth [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--service` | string | Migration service: todoist, trello, microsoft-todo, ticktick (required) |

**API endpoint:** `GET /api/v1/migration/{service}/auth`

### migration migrate

Start a migration from a service.

```
vikunja migration migrate [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--service` | string | Migration service (required) |
| `--code` | string | Auth code from the migration service (required) |

**API endpoint:** `POST /api/v1/migration/{service}/migrate`

### migration status

Get migration status for a service.

```
vikunja migration status [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--service` | string | Migration service (required) |

**API endpoint:** `GET /api/v1/migration/{service}/status`

---

## notifications

Manage notifications.

### notifications list

List notifications.

```
vikunja notifications list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--page` | int | Page number |
| `--per-page` | int | Items per page |

**API endpoint:** `GET /api/v1/notifications`

### notifications read

Mark a notification as read.

```
vikunja notifications read [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Notification ID (required) |

**API endpoint:** `POST /api/v1/notifications/{id}`

### notifications read-all

Mark all notifications as read.

```
vikunja notifications read-all
```

**API endpoint:** `POST /api/v1/notifications`

---

## projects

Manage projects.

### projects create

Create a new project.

```
vikunja projects create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--title` | string | Project title |
| `--description` | string | Project description |
| `--identifier` | string | Project identifier |
| `--hex-color` | string | Hex color (e.g. #ff0000) |
| `--parent-project-id` | int | Parent project ID |

**API endpoint:** `PUT /api/v1/projects`

### projects get

Get a project by ID.

```
vikunja projects get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Project ID |

**API endpoint:** `GET /api/v1/projects/{id}`

### projects list

List all projects.

```
vikunja projects list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |
| `--search` | string | Search projects by title |

**API endpoint:** `GET /api/v1/projects`

### projects update

Update a project.

```
vikunja projects update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Project ID |
| `--title` | string | Project title |
| `--description` | string | Project description |
| `--identifier` | string | Project identifier |
| `--hex-color` | string | Hex color (e.g. #ff0000) |
| `--is-archived` | bool | Archive the project |
| `--is-favorite` | bool | Mark as favorite |

**API endpoint:** `POST /api/v1/projects/{id}`

### projects delete

Delete a project.

```
vikunja projects delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Project ID |

**API endpoint:** `DELETE /api/v1/projects/{id}`

### projects duplicate

Duplicate a project.

```
vikunja projects duplicate [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Project ID to duplicate |

**API endpoint:** `PUT /api/v1/projects/{id}/duplicate`

### projects backgrounds

Manage project backgrounds.

#### projects backgrounds get

Download a project background.

```
vikunja projects backgrounds get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--output` | string | Output file path |

**API endpoint:** `GET /api/v1/projects/{id}/background`

#### projects backgrounds upload

Upload a project background image.

```
vikunja projects backgrounds upload [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--file` | string | Path to the background image file |

**API endpoint:** `PUT /api/v1/projects/{id}/backgrounds/upload` (multipart upload)

#### projects backgrounds delete

Delete a project background.

```
vikunja projects backgrounds delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |

**API endpoint:** `DELETE /api/v1/projects/{id}/background`

#### projects backgrounds search

Search Unsplash for background images.

```
vikunja projects backgrounds search [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--search` | string | Search query |
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |

**API endpoint:** `GET /api/v1/backgrounds/unsplash/search`

#### projects backgrounds unsplash

Set an Unsplash image as project background.

```
vikunja projects backgrounds unsplash [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--image-id` | string | Unsplash image ID |

**API endpoint:** `POST /api/v1/projects/{id}/backgrounds/unsplash`

### projects buckets

Manage project view buckets.

#### projects buckets create

Create a bucket in a project view.

```
vikunja projects buckets create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--view-id` | int | View ID |
| `--title` | string | Bucket title |
| `--limit` | int | Maximum number of tasks in this bucket |

**API endpoint:** `PUT /api/v1/projects/{projectID}/views/{viewID}/buckets`

#### projects buckets list

List buckets for a project view.

```
vikunja projects buckets list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--view-id` | int | View ID |
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |

**API endpoint:** `GET /api/v1/projects/{projectID}/views/{viewID}/buckets`

#### projects buckets update

Update a bucket in a project view.

```
vikunja projects buckets update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--view-id` | int | View ID |
| `--id` | int | Bucket ID |
| `--title` | string | Bucket title |
| `--limit` | int | Maximum number of tasks in this bucket |
| `--position` | float | Bucket position |

**API endpoint:** `POST /api/v1/projects/{projectID}/views/{viewID}/buckets/{id}`

#### projects buckets delete

Delete a bucket from a project view.

```
vikunja projects buckets delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--view-id` | int | View ID |
| `--id` | int | Bucket ID |

**API endpoint:** `DELETE /api/v1/projects/{projectID}/views/{viewID}/buckets/{id}`

### projects shares

Manage project shares.

#### projects shares create

Create a share link for a project.

```
vikunja projects shares create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--right` | int | Share right (0=read, 1=read&write, 2=admin) |
| `--sharing-type` | int | Sharing type (0=without password, 1=with password) |

**API endpoint:** `PUT /api/v1/projects/{id}/shares`

#### projects shares get

Get a share by hash.

```
vikunja projects shares get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--share-hash` | string | Share hash |

**API endpoint:** `GET /api/v1/shares/{hash}`

#### projects shares list

List shares for a project.

```
vikunja projects shares list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |

**API endpoint:** `GET /api/v1/projects/{id}/shares`

#### projects shares delete

Delete a project share.

```
vikunja projects shares delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--id` | int | Share ID |

**API endpoint:** `DELETE /api/v1/projects/{projectID}/shares/{id}`

#### projects shares auth

Authenticate against a share.

```
vikunja projects shares auth [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--share-hash` | string | Share hash |
| `--password` | string | Share password |

**API endpoint:** `POST /api/v1/shares/{hash}/auth`

### projects teams

Manage project team members.

#### projects teams add

Add a team to a project.

```
vikunja projects teams add [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--team-id` | int | Team ID |
| `--right` | int | Right (0=read, 1=read&write, 2=admin) |

**API endpoint:** `PUT /api/v1/projects/{projectID}/teams`

#### projects teams list

List teams assigned to a project.

```
vikunja projects teams list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |

**API endpoint:** `GET /api/v1/projects/{projectID}/teams`

#### projects teams update

Update a team's right on a project.

```
vikunja projects teams update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--team-id` | int | Team ID |
| `--right` | int | Right (0=read, 1=read&write, 2=admin) |

**API endpoint:** `POST /api/v1/projects/{projectID}/teams/{teamID}`

#### projects teams remove

Remove a team from a project.

```
vikunja projects teams remove [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--team-id` | int | Team ID |

**API endpoint:** `DELETE /api/v1/projects/{projectID}/teams/{teamID}`

### projects users

Manage project user members.

#### projects users add

Add a user to a project.

```
vikunja projects users add [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--user-id` | int | User ID |
| `--right` | int | Right (0=read, 1=read&write, 2=admin) |

**API endpoint:** `PUT /api/v1/projects/{projectID}/members`

#### projects users list

List user members of a project.

```
vikunja projects users list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |

**API endpoint:** `GET /api/v1/projects/{projectID}/members`

#### projects users update

Update a user's right on a project.

```
vikunja projects users update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--user-id` | int | User ID |
| `--right` | int | Right (0=read, 1=read&write, 2=admin) |

**API endpoint:** `POST /api/v1/projects/{projectID}/members/{userID}`

#### projects users remove

Remove a user from a project.

```
vikunja projects users remove [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--user-id` | int | User ID |

**API endpoint:** `DELETE /api/v1/projects/{projectID}/members/{userID}`

### projects views

Manage project views.

#### projects views create

Create a view for a project.

```
vikunja projects views create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--title` | string | View title |
| `--view-kind` | string | View kind (list, gantt, table, kanban) |
| `--filter` | string | View filter |
| `--bucket-config-mode` | string | Bucket configuration mode |

**API endpoint:** `PUT /api/v1/projects/{projectID}/views`

#### projects views get

Get a project view by ID.

```
vikunja projects views get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--id` | int | View ID |

**API endpoint:** `GET /api/v1/projects/{projectID}/views/{id}`

#### projects views list

List views for a project.

```
vikunja projects views list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |

**API endpoint:** `GET /api/v1/projects/{projectID}/views`

#### projects views update

Update a project view.

```
vikunja projects views update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--id` | int | View ID |
| `--title` | string | View title |
| `--filter` | string | View filter |

**API endpoint:** `POST /api/v1/projects/{projectID}/views/{id}`

#### projects views delete

Delete a project view.

```
vikunja projects views delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--id` | int | View ID |

**API endpoint:** `DELETE /api/v1/projects/{projectID}/views/{id}`

### projects webhooks

Manage project webhooks.

#### projects webhooks create

Create a webhook for a project.

```
vikunja projects webhooks create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--target-url` | string | Webhook target URL |
| `--events` | string | Comma-separated list of events |
| `--secret` | string | Webhook secret |

**API endpoint:** `PUT /api/v1/projects/{projectID}/webhooks`

#### projects webhooks list

List webhooks for a project.

```
vikunja projects webhooks list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--page` | int | Page number |
| `--per-page` | int | Number of items per page |

**API endpoint:** `GET /api/v1/projects/{projectID}/webhooks`

#### projects webhooks update

Update a project webhook.

```
vikunja projects webhooks update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--id` | int | Webhook ID |
| `--target-url` | string | Webhook target URL |
| `--events` | string | Comma-separated list of events |
| `--secret` | string | Webhook secret |

**API endpoint:** `POST /api/v1/projects/{projectID}/webhooks/{id}`

#### projects webhooks delete

Delete a project webhook.

```
vikunja projects webhooks delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID |
| `--id` | int | Webhook ID |

**API endpoint:** `DELETE /api/v1/projects/{projectID}/webhooks/{id}`

#### projects webhooks events

List available webhook events.

```
vikunja projects webhooks events
```

**API endpoint:** `GET /api/v1/webhooks/events`

---

## reactions

Manage reactions on tasks and comments.

### reactions add

Add a reaction to a task or comment.

```
vikunja reactions add [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--value` | string | Reaction value/emoji (required) |
| `--comment-id` | int | Comment ID (for comment reactions) |

**API endpoint:** `POST /api/v1/tasks/{taskID}/reactions` or `POST /api/v1/tasks/{taskID}/comments/{commentID}/reactions`

### reactions list

List reactions on a task or comment.

```
vikunja reactions list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--comment-id` | int | Comment ID (for comment reactions) |

**API endpoint:** `GET /api/v1/tasks/{taskID}/reactions` or `GET /api/v1/tasks/{taskID}/comments/{commentID}/reactions`

### reactions delete

Remove a reaction from a task or comment.

```
vikunja reactions delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--value` | string | Reaction value/emoji (required) |
| `--comment-id` | int | Comment ID (for comment reactions) |

**API endpoint:** `DELETE /api/v1/tasks/{taskID}/reactions` or `DELETE /api/v1/tasks/{taskID}/comments/{commentID}/reactions`

---

## subscriptions

Manage entity subscriptions. Supported entity types: `project`, `task`.

### subscriptions create

Subscribe to an entity.

```
vikunja subscriptions create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--entity-type` | string | Entity type: project or task (required) |
| `--entity-id` | int | Entity ID (required) |

**API endpoint:** `PUT /api/v1/{entityType}/{entityID}/subscription`

### subscriptions delete

Unsubscribe from an entity.

```
vikunja subscriptions delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--entity-type` | string | Entity type: project or task (required) |
| `--entity-id` | int | Entity ID (required) |

**API endpoint:** `DELETE /api/v1/{entityType}/{entityID}/subscription`

---

## system

System information.

### system info

Get system info from the Vikunja instance.

```
vikunja system info
```

**API endpoint:** `GET /api/v1/info`

---

## tasks

Manage tasks.

### tasks create

Create a new task in a project.

```
vikunja tasks create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--project-id` | int | Project ID (required) |
| `--title` | string | Task title (required) |
| `--description` | string | Task description |
| `--done` | bool | Mark task as done |
| `--priority` | int | Priority (0-5) |
| `--due-date` | string | Due date (RFC3339) |
| `--start-date` | string | Start date (RFC3339) |
| `--end-date` | string | End date (RFC3339) |
| `--hex-color` | string | Hex color code |
| `--percent-done` | float | Percent done (0.0-1.0) |
| `--repeat-after` | int | Repeat after N seconds |
| `--is-favorite` | bool | Mark as favorite |
| `--bucket-id` | int | Bucket ID |

**API endpoint:** `PUT /api/v1/projects/{projectID}/tasks`

### tasks get

Get a task by ID.

```
vikunja tasks get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Task ID (required) |

**API endpoint:** `GET /api/v1/tasks/{id}`

### tasks list

List all tasks across all projects.

```
vikunja tasks list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--page` | int | Page number |
| `--per-page` | int | Items per page |
| `--search` | string | Search tasks |
| `--sort` | string | Sort by field |
| `--order-by` | string | Order direction (asc/desc) |
| `--filter` | string | Filter query |
| `--filter-by` | string | Field to filter by |
| `--filter-value` | string | Value to filter for |
| `--filter-comparator` | string | Filter comparator (equals, greater, less, etc.) |

**API endpoint:** `GET /api/v1/tasks/all`

### tasks update

Update a task.

```
vikunja tasks update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Task ID (required) |
| `--title` | string | Task title |
| `--description` | string | Task description |
| `--done` | bool | Mark task as done |
| `--priority` | int | Priority (0-5) |
| `--due-date` | string | Due date (RFC3339) |
| `--start-date` | string | Start date (RFC3339) |
| `--end-date` | string | End date (RFC3339) |
| `--hex-color` | string | Hex color code |
| `--percent-done` | float | Percent done (0.0-1.0) |
| `--repeat-after` | int | Repeat after N seconds |
| `--is-favorite` | bool | Mark as favorite |
| `--bucket-id` | int | Bucket ID |

**API endpoint:** `POST /api/v1/tasks/{id}`

### tasks delete

Delete a task.

```
vikunja tasks delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Task ID (required) |

**API endpoint:** `DELETE /api/v1/tasks/{id}`

### tasks bulk

Bulk update tasks.

```
vikunja tasks bulk [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-ids` | string | Comma-separated task IDs (required) |
| `--done` | bool | Mark tasks as done |
| `--priority` | int | Priority (0-5) |
| `--due-date` | string | Due date (RFC3339) |

**API endpoint:** `POST /api/v1/tasks/bulk`

### tasks assignees

Manage task assignees.

#### tasks assignees add

Add an assignee to a task.

```
vikunja tasks assignees add [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--user-id` | int | User ID to assign (required) |

**API endpoint:** `PUT /api/v1/tasks/{taskID}/assignees`

#### tasks assignees list

List assignees for a task.

```
vikunja tasks assignees list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--page` | int | Page number |
| `--per-page` | int | Items per page |

**API endpoint:** `GET /api/v1/tasks/{taskID}/assignees`

#### tasks assignees remove

Remove an assignee from a task.

```
vikunja tasks assignees remove [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--user-id` | int | User ID to remove (required) |

**API endpoint:** `DELETE /api/v1/tasks/{taskID}/assignees/{userID}`

#### tasks assignees bulk

Bulk assign users to a task.

```
vikunja tasks assignees bulk [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--user-ids` | string | Comma-separated user IDs (required) |

**API endpoint:** `POST /api/v1/tasks/{taskID}/assignees/bulk`

### tasks attachments

Manage task attachments.

#### tasks attachments upload

Upload an attachment to a task.

```
vikunja tasks attachments upload [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--file` | string | Path to the file to upload (required) |

**API endpoint:** `PUT /api/v1/tasks/{taskID}/attachments` (multipart upload)

#### tasks attachments list

List attachments for a task.

```
vikunja tasks attachments list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--page` | int | Page number |
| `--per-page` | int | Items per page |

**API endpoint:** `GET /api/v1/tasks/{taskID}/attachments`

#### tasks attachments get

Download an attachment.

```
vikunja tasks attachments get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--id` | int | Attachment ID (required) |
| `--output` | string | Output file path (required) |

**API endpoint:** `GET /api/v1/tasks/{taskID}/attachments/{id}`

#### tasks attachments delete

Delete an attachment.

```
vikunja tasks attachments delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--id` | int | Attachment ID (required) |

**API endpoint:** `DELETE /api/v1/tasks/{taskID}/attachments/{id}`

### tasks comments

Manage task comments.

#### tasks comments create

Create a comment on a task.

```
vikunja tasks comments create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--comment` | string | Comment text (required) |

**API endpoint:** `PUT /api/v1/tasks/{taskID}/comments`

#### tasks comments get

Get a comment by ID.

```
vikunja tasks comments get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--id` | int | Comment ID (required) |

**API endpoint:** `GET /api/v1/tasks/{taskID}/comments/{id}`

#### tasks comments list

List comments for a task.

```
vikunja tasks comments list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--page` | int | Page number |
| `--per-page` | int | Items per page |

**API endpoint:** `GET /api/v1/tasks/{taskID}/comments`

#### tasks comments update

Update a comment.

```
vikunja tasks comments update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--id` | int | Comment ID (required) |
| `--comment` | string | Updated comment text (required) |

**API endpoint:** `POST /api/v1/tasks/{taskID}/comments/{id}`

#### tasks comments delete

Delete a comment.

```
vikunja tasks comments delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--id` | int | Comment ID (required) |

**API endpoint:** `DELETE /api/v1/tasks/{taskID}/comments/{id}`

### tasks labels

Manage task labels.

#### tasks labels add

Add a label to a task.

```
vikunja tasks labels add [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--label-id` | int | Label ID to add (required) |

**API endpoint:** `PUT /api/v1/tasks/{taskID}/labels`

#### tasks labels list

List labels for a task.

```
vikunja tasks labels list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--page` | int | Page number |
| `--per-page` | int | Items per page |

**API endpoint:** `GET /api/v1/tasks/{taskID}/labels`

#### tasks labels remove

Remove a label from a task.

```
vikunja tasks labels remove [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--label-id` | int | Label ID to remove (required) |

**API endpoint:** `DELETE /api/v1/tasks/{taskID}/labels/{labelID}`

#### tasks labels bulk

Bulk assign labels to a task.

```
vikunja tasks labels bulk [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--label-ids` | string | Comma-separated label IDs (required) |

**API endpoint:** `POST /api/v1/tasks/{taskID}/labels/bulk`

### tasks relations

Manage task relations.

#### tasks relations create

Create a relation between two tasks.

```
vikunja tasks relations create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--other-task-id` | int | Other task ID (required) |
| `--kind` | string | Relation kind (required) |

**API endpoint:** `PUT /api/v1/tasks/{taskID}/relations`

#### tasks relations delete

Delete a relation between two tasks.

```
vikunja tasks relations delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--task-id` | int | Task ID (required) |
| `--other-task-id` | int | Other task ID (required) |
| `--kind` | string | Relation kind (required) |

**API endpoint:** `DELETE /api/v1/tasks/{taskID}/relations`

---

## teams

Manage teams.

### teams create

Create a team.

```
vikunja teams create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--name` | string | Team name (required) |
| `--description` | string | Team description |

**API endpoint:** `PUT /api/v1/teams`

### teams get

Get a team by ID.

```
vikunja teams get [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Team ID (required) |

**API endpoint:** `GET /api/v1/teams/{id}`

### teams list

List all teams.

```
vikunja teams list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--page` | int | Page number |
| `--per-page` | int | Items per page |
| `--search` | string | Search teams |

**API endpoint:** `GET /api/v1/teams`

### teams update

Update a team.

```
vikunja teams update [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Team ID (required) |
| `--name` | string | Team name |
| `--description` | string | Team description |

**API endpoint:** `POST /api/v1/teams/{id}`

### teams delete

Delete a team.

```
vikunja teams delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Team ID (required) |

**API endpoint:** `DELETE /api/v1/teams/{id}`

### teams members

Manage team members.

#### teams members add

Add a member to a team.

```
vikunja teams members add [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--team-id` | int | Team ID (required) |
| `--user-id` | int | User ID (required) |

**API endpoint:** `PUT /api/v1/teams/{teamID}/members`

#### teams members remove

Remove a member from a team.

```
vikunja teams members remove [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--team-id` | int | Team ID (required) |
| `--user-id` | int | User ID (required) |

**API endpoint:** `DELETE /api/v1/teams/{teamID}/members/{userID}`

#### teams members admin

Toggle admin status of a team member.

```
vikunja teams members admin [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--team-id` | int | Team ID (required) |
| `--user-id` | int | User ID (required) |
| `--admin` | bool | Set admin status |

**API endpoint:** `POST /api/v1/teams/{teamID}/members/{userID}`

---

## tokens

Manage API tokens.

### tokens create

Create an API token.

```
vikunja tokens create [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--title` | string | Token title (required) |
| `--permissions` | string | Permissions as JSON |
| `--expires-at` | string | Expiration date (RFC3339) |

**API endpoint:** `PUT /api/v1/tokens`

### tokens list

List all API tokens.

```
vikunja tokens list
```

**API endpoint:** `GET /api/v1/tokens`

### tokens delete

Delete an API token.

```
vikunja tokens delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Token ID (required) |

**API endpoint:** `DELETE /api/v1/tokens/{id}`

### tokens test

Test the current API token.

```
vikunja tokens test
```

**API endpoint:** `GET /api/v1/tokenTest`

### tokens routes

List available API routes for token permissions.

```
vikunja tokens routes
```

**API endpoint:** `GET /api/v1/routes`

---

## user

Manage user account.

### user get

Get the current user.

```
vikunja user get
```

**API endpoint:** `GET /api/v1/user`

### user list

Search/list users.

```
vikunja user list [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--search` | string | Search query (required) |
| `--page` | int | Page number |
| `--per-page` | int | Items per page |

**API endpoint:** `GET /api/v1/users`

### user password

Change user password.

```
vikunja user password [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--old-password` | string | Current password (required) |
| `--new-password` | string | New password (required) |

**API endpoint:** `POST /api/v1/user/password`

### user token

Get a new JWT token for the current user.

```
vikunja user token
```

**API endpoint:** `POST /api/v1/user/token`

### user caldav-tokens

Manage CalDAV tokens.

#### user caldav-tokens create

Create a CalDAV token.

```
vikunja user caldav-tokens create
```

**API endpoint:** `PUT /api/v1/user/settings/token/caldav`

#### user caldav-tokens list

List CalDAV tokens.

```
vikunja user caldav-tokens list
```

**API endpoint:** `GET /api/v1/user/settings/token/caldav`

#### user caldav-tokens delete

Delete a CalDAV token.

```
vikunja user caldav-tokens delete [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--id` | int | Token ID (required) |

**API endpoint:** `DELETE /api/v1/user/settings/token/caldav/{id}`

### user deletion

Manage account deletion.

#### user deletion request

Request account deletion.

```
vikunja user deletion request [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--password` | string | Account password (required) |

**API endpoint:** `POST /api/v1/user/deletion/request`

#### user deletion confirm

Confirm account deletion.

```
vikunja user deletion confirm [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--token` | string | Confirmation token (required) |

**API endpoint:** `POST /api/v1/user/deletion/confirm`

#### user deletion cancel

Cancel account deletion.

```
vikunja user deletion cancel [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--password` | string | Account password (required) |

**API endpoint:** `POST /api/v1/user/deletion/cancel`

### user export

Manage data export.

#### user export request

Request a data export.

```
vikunja user export request [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--password` | string | Account password (required) |

**API endpoint:** `POST /api/v1/user/export/request`

#### user export status

Check export status.

```
vikunja user export status
```

**API endpoint:** `GET /api/v1/user/export/status`

#### user export download

Download the data export.

```
vikunja user export download [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--output` | string | Output file path (default "vikunja-export.zip") |

**API endpoint:** `GET /api/v1/user/export/download`

### user settings

Manage user settings.

#### user settings general

Get or update general settings. Without flags, returns current settings. With flags, updates them.

```
vikunja user settings general [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--name` | string | Display name |
| `--language` | string | Language code |
| `--timezone` | string | Timezone |
| `--week-start` | int | Week start day (0=Sunday, 1=Monday) |
| `--default-project-id` | int | Default project ID |
| `--discover-by-email` | bool | Discoverable by email |
| `--discover-by-name` | bool | Discoverable by name |
| `--email-reminders` | bool | Enable email reminders |
| `--overdue-reminders` | bool | Enable overdue task reminders |
| `--frontend-settings` | string | Frontend settings as JSON |

**API endpoint:** `GET /api/v1/user/settings/general` (read) or `POST /api/v1/user/settings/general` (update)

#### user settings email

Update user email address.

```
vikunja user settings email [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--email` | string | New email address (required) |

**API endpoint:** `POST /api/v1/user/settings/email`

#### user settings avatar

Set avatar provider or upload avatar.

```
vikunja user settings avatar [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--provider` | string | Avatar provider: gravatar, initials, upload, marble (required) |

**API endpoint:** `POST /api/v1/user/settings/avatar`

#### user settings timezones

List available timezones.

```
vikunja user settings timezones
```

**API endpoint:** `GET /api/v1/user/timezones`

### user totp

Manage TOTP two-factor authentication.

#### user totp enroll

Enroll in TOTP (get secret and URL).

```
vikunja user totp enroll
```

**API endpoint:** `POST /api/v1/user/settings/totp/enroll`

#### user totp enable

Enable TOTP with a passcode.

```
vikunja user totp enable [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--passcode` | string | TOTP passcode (required) |

**API endpoint:** `POST /api/v1/user/settings/totp/enable`

#### user totp disable

Disable TOTP.

```
vikunja user totp disable [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--password` | string | Account password (required) |

**API endpoint:** `POST /api/v1/user/settings/totp/disable`

#### user totp status

Get TOTP status.

```
vikunja user totp status
```

**API endpoint:** `GET /api/v1/user/settings/totp`

#### user totp qrcode

Download TOTP QR code image.

```
vikunja user totp qrcode [flags]
```

| Flag | Type | Description |
|---|---|---|
| `--output` | string | Output file path (default "totp-qrcode.png") |

**API endpoint:** `GET /api/v1/user/settings/totp/qrcode`

---

## version

Print the CLI version, commit hash, and build date.

```
vikunja version
```

This command does not call any API endpoint.

---

## Shell Completion

Generate shell completion scripts for tab-completion of commands and flags.

### Bash

```bash
# Load in current session
source <(vikunja completion bash)

# Install permanently (Linux)
vikunja completion bash > /etc/bash_completion.d/vikunja

# Install permanently (macOS)
vikunja completion bash > $(brew --prefix)/etc/bash_completion.d/vikunja
```

Requires the `bash-completion` package.

| Flag | Type | Description |
|---|---|---|
| `--no-descriptions` | bool | Disable completion descriptions |

### Zsh

```bash
# Load in current session
source <(vikunja completion zsh)

# Install permanently (Linux)
vikunja completion zsh > "${fpath[1]}/_vikunja"

# Install permanently (macOS)
vikunja completion zsh > $(brew --prefix)/share/zsh/site-functions/_vikunja
```

You may need to enable shell completion first:

```bash
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

| Flag | Type | Description |
|---|---|---|
| `--no-descriptions` | bool | Disable completion descriptions |

### Fish

```bash
# Load in current session
vikunja completion fish | source

# Install permanently
vikunja completion fish > ~/.config/fish/completions/vikunja.fish
```

| Flag | Type | Description |
|---|---|---|
| `--no-descriptions` | bool | Disable completion descriptions |

### PowerShell

```powershell
# Load in current session
vikunja completion powershell | Out-String | Invoke-Expression
```

To load for every session, add the output of the above command to your PowerShell profile.

| Flag | Type | Description |
|---|---|---|
| `--no-descriptions` | bool | Disable completion descriptions |
