// github_client.go
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mwelwankuta/unethical-commits/models"
)

type GitHubClient struct {
	BaseURL    string
	AuthToken  string
	HttpClient *http.Client
}

func NewGitHubClient(baseURL, authToken string) *GitHubClient {
	return &GitHubClient{
		BaseURL:    baseURL,
		AuthToken:  authToken,
		HttpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *GitHubClient) createRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	url := c.BaseURL + endpoint
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	return req, nil
}

func (c *GitHubClient) CreateContribution(owner, repo string, issue models.GitHubIssue) error {
	endpoint := fmt.Sprintf("/repos/%s/%s/issues", owner, repo)
	_, err := c.createRequest("POST", c.BaseURL+endpoint, issue)
	if err != nil {
		return err
	}

	fmt.Println(c.BaseURL + endpoint)

	return nil
}
