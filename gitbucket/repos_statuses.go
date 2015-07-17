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
