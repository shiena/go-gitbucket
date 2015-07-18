package gitbucket

import (
	"fmt"
	"net/http"
	"time"
)

type RepoStatus struct {
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	State       *string    `json:"state"`
	TargetURL   *string    `json:"target_url,omitempty"`
	Description *string    `json:"description,omitempty"`
	ID          *int       `json:"id,omitempty"`
	Context     *string    `json:"context,omitempty"`
	Creator     *User      `json:"creator,omitempty"`
	URL         *string    `json:"url,omitempty"`
}

func (s *RepositoriesService) CreateStatus(owner, repo, sha string, status *RepoStatus) (*RepoStatus, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/statuses/%v", owner, repo, sha)
	req, err := s.client.NewRequest("POST", u, status)
	if err != nil {
		return nil, nil, err
	}

	repoStatus := new(RepoStatus)
	resp, err := s.client.Do(req, repoStatus)
	if err != nil {
		return nil, resp, err
	}

	return repoStatus, resp, err
}

func (s *RepositoriesService) ListStatuses(owner, repo, ref string) ([]RepoStatus, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/commits/%v/statuses", owner, repo, ref)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	statuses := new([]RepoStatus)
	resp, err := s.client.Do(req, statuses)
	if err != nil {
		return nil, resp, err
	}

	return *statuses, resp, err
}

type CombinedStatus struct {
	State      *string      `json:"state"`
	SHA        *string      `json:"sha"`
	TotalCount *int         `json:"total_count"`
	Statuses   []RepoStatus `json:"statuses"`
	Repository *Repository  `json:"repository"`
	URL        *string      `json:"url"`
}

func (s *RepositoriesService) GetCombinedStatus(owner, repo, ref string) (*CombinedStatus, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/commits/%v/status", owner, repo, ref)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(CombinedStatus)
	resp, err := s.client.Do(req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, err
}
