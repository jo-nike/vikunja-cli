package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/jo-nike/vikunja-cli/internal/resolve"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var page, perPage int
	var search, sort, orderBy string
	var filter, filterBy, filterValue, filterComparator string
	var projectID, bucketID, viewID int64
	var bucket string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks across all projects",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

			hasBucket := cmd.Flags().Changed("bucket")
			hasBucketID := cmd.Flags().Changed("bucket-id")

			if hasBucket && hasBucketID {
				output.Error(fmt.Errorf("--bucket and --bucket-id are mutually exclusive"))
			}

			// Bucket filtering: use kanban view tasks endpoint
			if hasBucket || hasBucketID {
				if projectID == 0 {
					output.Error(fmt.Errorf("--project-id is required when using --bucket or --bucket-id"))
				}

				// Resolve view ID
				if viewID == 0 {
					view, err := resolve.FindKanbanView(c, projectID)
					if err != nil {
						output.Error(err)
					}
					viewID = view.ID
				}

				// Resolve bucket name to ID
				var targetBucketID int64
				if hasBucket {
					b, err := resolve.BucketByName(c, projectID, viewID, bucket)
					if err != nil {
						output.Error(err)
					}
					targetBucketID = b.ID
				} else {
					targetBucketID = bucketID
				}

				// Build request options for the view tasks endpoint
				var opts []client.RequestOption
				if page > 0 {
					opts = append(opts, client.WithPage(page))
				}
				if perPage > 0 {
					opts = append(opts, client.WithPerPage(perPage))
				}
				if search != "" {
					opts = append(opts, client.WithSearch(search))
				}
				if sort != "" {
					opts = append(opts, client.WithSort(sort))
				}
				if orderBy != "" {
					opts = append(opts, client.WithOrderBy(orderBy))
				}
				if filter != "" {
					opts = append(opts, client.WithFilter(filter))
				}
				if filterBy != "" {
					opts = append(opts, client.WithFilterBy(filterBy))
				}
				if filterValue != "" {
					opts = append(opts, client.WithFilterValue(filterValue))
				}
				if filterComparator != "" {
					opts = append(opts, client.WithFilterComparator(filterComparator))
				}

				buckets, _, err := resolve.ViewTaskBuckets(c, projectID, viewID, opts...)
				if err != nil {
					output.Error(err)
				}

				// Find the target bucket and stamp bucket_id on each task
				var tasks []models.Task
				for _, b := range buckets {
					if b.ID == targetBucketID {
						for i := range b.Tasks {
							b.Tasks[i].BucketID = b.ID
						}
						tasks = b.Tasks
						break
					}
				}
				if tasks == nil {
					tasks = []models.Task{}
				}

				info := &client.PaginationInfo{
					Page:       1,
					TotalPages: 1,
					TotalItems: len(tasks),
					PerPage:    len(tasks),
				}
				output.ResultList(tasks, info)
				return
			}

			// Default behavior: list all tasks
			if cmd.Flags().Changed("project-id") {
				pf := fmt.Sprintf("project_id = %d", projectID)
				if filter != "" {
					filter = filter + " && " + pf
				} else {
					filter = pf
				}
			}

			var opts []client.RequestOption
			if page > 0 {
				opts = append(opts, client.WithPage(page))
			}
			if perPage > 0 {
				opts = append(opts, client.WithPerPage(perPage))
			}
			if search != "" {
				opts = append(opts, client.WithSearch(search))
			}
			if sort != "" {
				opts = append(opts, client.WithSort(sort))
			}
			if orderBy != "" {
				opts = append(opts, client.WithOrderBy(orderBy))
			}
			if filter != "" {
				opts = append(opts, client.WithFilter(filter))
			}
			if filterBy != "" {
				opts = append(opts, client.WithFilterBy(filterBy))
			}
			if filterValue != "" {
				opts = append(opts, client.WithFilterValue(filterValue))
			}
			if filterComparator != "" {
				opts = append(opts, client.WithFilterComparator(filterComparator))
			}

			var result []models.Task
			info, err := c.GetList("/tasks", &result, opts...)
			if err != nil {
				output.Error(err)
			}
			output.ResultList(result, info)
		},
	}
	cmd.Flags().IntVar(&page, "page", 0, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 0, "Items per page")
	cmd.Flags().StringVar(&search, "search", "", "Search tasks")
	cmd.Flags().StringVar(&sort, "sort", "", "Sort by field")
	cmd.Flags().StringVar(&orderBy, "order-by", "", "Order direction (asc/desc)")
	cmd.Flags().StringVar(&filter, "filter", "", "Filter query")
	cmd.Flags().StringVar(&filterBy, "filter-by", "", "Field to filter by")
	cmd.Flags().StringVar(&filterValue, "filter-value", "", "Value to filter for")
	cmd.Flags().StringVar(&filterComparator, "filter-comparator", "", "Filter comparator (equals, greater, less, etc.)")
	cmd.Flags().Int64Var(&projectID, "project-id", 0, "Filter tasks by project ID")
	cmd.Flags().StringVar(&bucket, "bucket", "", "Filter by bucket name (requires --project-id)")
	cmd.Flags().Int64Var(&bucketID, "bucket-id", 0, "Filter by bucket ID (requires --project-id)")
	cmd.Flags().Int64Var(&viewID, "view-id", 0, "Kanban view ID (auto-detected if not specified)")
	return cmd
}
