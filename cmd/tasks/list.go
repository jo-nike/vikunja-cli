package tasks

import (
	"fmt"

	"github.com/jo-nike/vikunja-cli/internal/client"
	"github.com/jo-nike/vikunja-cli/internal/cmdutil"
	"github.com/jo-nike/vikunja-cli/internal/models"
	"github.com/jo-nike/vikunja-cli/internal/output"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var page, perPage int
	var search, sort, orderBy string
	var filter, filterBy, filterValue, filterComparator string
	var projectID int64

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks across all projects",
		Run: func(cmd *cobra.Command, args []string) {
			c := cmdutil.MustClient()

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
	return cmd
}
