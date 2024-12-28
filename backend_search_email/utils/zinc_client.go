package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	zincURL  = os.Getenv("ZINC_URL")
	username = os.Getenv("ZINC_USERNAME")
	password = os.Getenv("ZINC_PASSWORD")
)

func QueryZincSearch(index string, params map[string]string) ([]map[string]interface{}, error) {
	url := zincURL + "/" + index + "/_search"
	query := params["query"]
	messageID := params["message_id"]
	page := 0 
	size := 20 
	if p, ok := params["page"]; ok {
		var err error
		page, err = strconv.Atoi(p)
		if err != nil {
			page = 0 
		}
	}
	if s, ok := params["size"]; ok {
		var err error
		size, err = strconv.Atoi(s)
		if err != nil {
			size = 20 
		}
	}

	from := page * size 

	var queryBody map[string]interface{}
	if messageID != "" {
		queryBody = map[string]interface{}{
			"search_type": "querystring",
			"query": map[string]interface{}{
				"term":  messageID,
				"field": "message_id",
			},
			"from":        from,
			"size":        size,
			"_source":     []string{"from", "to", "content", "subject", "message_id"},
		}
	} else if query != "" {
		queryBody = map[string]interface{}{
			"search_type": "match",
			"query": map[string]interface{}{
				"term":  query,
				"field": "_all",
			},
			"from":        from,
			"size":        size,
			"_source":     []string{"from", "to", "content", "subject", "message_id"},
		}
	} else {
		queryBody = map[string]interface{}{
			"search_type": "match_all",
			"from":        from,
			"size":        size,
			"_source":     []string{"from", "to", "content", "subject", "message_id"},
		}
	}

	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("ZincSearch query failed: " + resp.Status)
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source map[string]interface{} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var documents []map[string]interface{}
	for _, hit := range result.Hits.Hits {
		documents = append(documents, hit.Source)
	}

	return documents, nil
}
