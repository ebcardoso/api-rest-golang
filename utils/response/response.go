package response

import (
	"encoding/json"
	"net/http"
)

func JsonRes(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
