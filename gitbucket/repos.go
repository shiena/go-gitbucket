package gitbucket

import (
	"fmt"
	"net/http"
)

type RepositoriesService struct {
	client *Client
}

// Repository represents a API user.
type Repository struct {
	Name          *string `json:"name"`
	FullName      *string `json:"full_name"`
	Description   *string `json:"description"`
	Watchers      *int    `json:"watchers"`
	Forks         *int    `json:"forks"`
	Private       *bool   `json:"private"`
	DefaultBranch *string `json:"default_branch"`
	Owner         *User   `json:"owner"`
	ForksCount    *int    `json:"forks_count"`
	WatchersCount *int    `json:"watchers_count"`
	URL           *string `json:"url"`
	HTTPURL       *string `json:"http_url"`
	CloneURL      *string `json:"clone_url"`
	HTMLURL       *string `json:"html_url"`
}

func (s *RepositoriesService) Get(owner, repo string) (*Repository, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Repository)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}
