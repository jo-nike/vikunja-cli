package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jo-nike/vikunja-cli/internal/client"
)

func Result(data interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

type PaginatedResponse struct {
	Data        interface{} `json:"data"`
	Page        int         `json:"page"`
	TotalPages  int         `json:"total_pages"`
	ResultCount int         `json:"result_count"`
}

func ResultList(data interface{}, info *client.PaginationInfo) {
	resp := PaginatedResponse{
		Data:        data,
		Page:        info.Page,
		TotalPages:  info.TotalPages,
		ResultCount: info.TotalItems,
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(resp)
}

func Error(err error) {
	msg := map[string]string{"error": err.Error()}
	data, _ := json.MarshalIndent(msg, "", "  ")
	fmt.Fprintln(os.Stderr, string(data))
	os.Exit(1)
}
