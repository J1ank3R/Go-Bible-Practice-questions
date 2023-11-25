package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func Parse(data *IssuesSearchResult) map[string][]*Issue {
	hash := make(map[string][]*Issue)

	hash["A_mouth"] = make([]*Issue, 0)
	hash["A_year"] = make([]*Issue, 0)
	hash["More_than_one_year"] = make([]*Issue, 0)

	var now = time.Now()

	for _, val := range data.Items {
		if val.CreateAt.Year() == now.Year() {
			if val.CreateAt.Month() == now.Month() {
				hash["A_mouth"] = append(hash["A_mouth"], val)
			} else {
				hash["A_year"] = append(hash["A_year"], val)
			}
		} else {
			hash["More_than_one_year"] = append(hash["More_than_one_year"], val)
		}
	}
	return hash
}
