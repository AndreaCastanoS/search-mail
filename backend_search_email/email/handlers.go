package email

import (
	"chellenge/utils"
	"encoding/json"
	"net/http"
	"strings"
)

type SearchRequest struct {
	Query string `json:"q"`
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	startTime := strings.TrimSpace(r.URL.Query().Get("start_time"))
	endTime := strings.TrimSpace(r.URL.Query().Get("end_time"))


	if len(query) > 100 {
		http.Error(w, "Query parameter 'q' too long", http.StatusBadRequest)
		return
	}

	indexName := "emails"

	params := map[string]string{
		"query":     query,
		"startTime": startTime,
		"endTime":   endTime,
	}

	results, err := utils.QueryZincSearch(indexName, params)
	if err != nil {
		http.Error(w, "Error querying ZincSearch: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		http.Error(w, "No results found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"results": results,
	})
}

func GetEmailByID(w http.ResponseWriter, r *http.Request) {
	messageID := r.URL.Query().Get("message_id")

	if messageID == "" {
		http.Error(w, "Parameter 'message_id' is required", http.StatusBadRequest)
		return
	}

	indexName := "emails"
	params := map[string]string{
		"query": messageID,
	}

	results, err := utils.QueryZincSearch(indexName, params)
	if err != nil {
		http.Error(w, "Error querying ZincSearch: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results[0])
}