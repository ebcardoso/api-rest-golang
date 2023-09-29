package response

import (
	"encoding/json"
	"net/http"
)

type jsonPaginatedResponse struct {
	Page          int64       `json:"page,omitempty"`
	PageSize      int64       `json:"page_size,omitempty"`
	TotalPages    int64       `json:"total_pages,omitempty"`
	TotalElements int64       `json:"total_elements,omitempty"`
	Content       interface{} `json:"content,omitempty"`
}

func JsonRes(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func JsonPaginatedResponse(w http.ResponseWriter, body interface{}, control []int64) {
	output := jsonPaginatedResponse{
		Page:          control[0],
		PageSize:      control[1],
		TotalPages:    (control[2] / 10) + 1,
		TotalElements: control[2],
		Content:       body,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
