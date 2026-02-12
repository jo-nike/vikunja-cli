package resolve

import (
	"fmt"
	"strings"

	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/models"
)

// FindKanbanView lists views for a project and returns the single kanban view.
// Returns an error if zero or multiple kanban views exist, listing them with IDs.
func FindKanbanView(c *client.Client, projectID int64) (*models.ProjectView, error) {
	path := fmt.Sprintf("/projects/%d/views", projectID)
	var views []models.ProjectView
	if _, err := c.GetList(path, &views); err != nil {
		return nil, fmt.Errorf("listing views for project %d: %w", projectID, err)
	}

	var kanbanViews []models.ProjectView
	for _, v := range views {
		if v.ViewKind == "kanban" {
			kanbanViews = append(kanbanViews, v)
		}
	}

	switch len(kanbanViews) {
	case 0:
		return nil, fmt.Errorf("no kanban view found for project %d; use --view-id to specify a view", projectID)
	case 1:
		return &kanbanViews[0], nil
	default:
		var parts []string
		for _, v := range kanbanViews {
			parts = append(parts, fmt.Sprintf("  %d: %s", v.ID, v.Title))
		}
		return nil, fmt.Errorf("multiple kanban views found for project %d; use --view-id to specify one:\n%s", projectID, strings.Join(parts, "\n"))
	}
}

// BucketByName performs a case-insensitive bucket name lookup within a view.
// Returns an error if not found (listing available buckets) or if ambiguous.
func BucketByName(c *client.Client, projectID, viewID int64, name string) (*models.Bucket, error) {
	path := fmt.Sprintf("/projects/%d/views/%d/buckets", projectID, viewID)
	var buckets []models.Bucket
	if _, err := c.GetList(path, &buckets); err != nil {
		return nil, fmt.Errorf("listing buckets for project %d view %d: %w", projectID, viewID, err)
	}

	nameLower := strings.ToLower(name)
	var matches []models.Bucket
	for _, b := range buckets {
		if strings.ToLower(b.Title) == nameLower {
			matches = append(matches, b)
		}
	}

	switch len(matches) {
	case 0:
		var available []string
		for _, b := range buckets {
			available = append(available, fmt.Sprintf("  %d: %s", b.ID, b.Title))
		}
		if len(available) == 0 {
			return nil, fmt.Errorf("no buckets found in project %d view %d", projectID, viewID)
		}
		return nil, fmt.Errorf("bucket %q not found; available buckets:\n%s", name, strings.Join(available, "\n"))
	case 1:
		return &matches[0], nil
	default:
		var parts []string
		for _, b := range matches {
			parts = append(parts, fmt.Sprintf("  %d: %s", b.ID, b.Title))
		}
		return nil, fmt.Errorf("multiple buckets match %q; use --bucket-id to specify one:\n%s", name, strings.Join(parts, "\n"))
	}
}

// BucketByNameAutoView resolves a bucket by name. If viewID is 0, it auto-detects the
// kanban view first via FindKanbanView, then calls BucketByName.
func BucketByNameAutoView(c *client.Client, projectID, viewID int64, name string) (*models.Bucket, int64, error) {
	if viewID == 0 {
		view, err := FindKanbanView(c, projectID)
		if err != nil {
			return nil, 0, err
		}
		viewID = view.ID
	}

	bucket, err := BucketByName(c, projectID, viewID, name)
	if err != nil {
		return nil, 0, err
	}
	return bucket, viewID, nil
}

// TaskProjectID fetches a task and returns its project_id.
func TaskProjectID(c *client.Client, taskID int64) (int64, error) {
	path := fmt.Sprintf("/tasks/%d", taskID)
	var task models.Task
	if err := c.Get(path, &task); err != nil {
		return 0, fmt.Errorf("fetching task %d: %w", taskID, err)
	}
	if task.ProjectID == 0 {
		return 0, fmt.Errorf("task %d has no project_id", taskID)
	}
	return task.ProjectID, nil
}

// ViewTaskBuckets fetches bucket-organized tasks from a kanban view.
// Returns []TaskBucket (buckets with embedded tasks) and pagination info.
func ViewTaskBuckets(c *client.Client, projectID, viewID int64, opts ...client.RequestOption) ([]models.TaskBucket, *client.PaginationInfo, error) {
	path := fmt.Sprintf("/projects/%d/views/%d/tasks", projectID, viewID)
	var buckets []models.TaskBucket
	info, err := c.GetList(path, &buckets, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("fetching tasks for project %d view %d: %w", projectID, viewID, err)
	}
	return buckets, info, nil
}

// TaskBucketID resolves the real bucket_id for a task by finding it in kanban view data.
// If projectID is 0, it fetches the task to get its project_id.
// If viewID is 0, it auto-detects the kanban view via FindKanbanView.
// Returns the bucket ID, or 0 if the task is not found in any bucket.
func TaskBucketID(c *client.Client, taskID, projectID, viewID int64) (int64, error) {
	if projectID == 0 {
		pid, err := TaskProjectID(c, taskID)
		if err != nil {
			return 0, err
		}
		projectID = pid
	}
	if viewID == 0 {
		view, err := FindKanbanView(c, projectID)
		if err != nil {
			return 0, err
		}
		viewID = view.ID
	}

	filterOpt := client.WithFilter(fmt.Sprintf("id = %d", taskID))
	buckets, _, err := ViewTaskBuckets(c, projectID, viewID, filterOpt)
	if err != nil {
		return 0, err
	}

	for _, b := range buckets {
		for _, t := range b.Tasks {
			if t.ID == taskID {
				return b.ID, nil
			}
		}
	}
	return 0, nil
}
