package gitbucket

import (
	"fmt"
	"net/http"
	"time"
)

type PullRequestsService struct {
	client *Client
}

type PullRequestBranch struct {
	SHA   *string     `json:"sha"`
	Ref   *string     `json:"ref"`
	Repo  *Repository `json:"repo"`
	Label *string     `json:"label"`
	User  *User       `json:"user"`
}

type PullRequest struct {
	Number            *int               `json:"number"`
	UpdatedAt         *time.Time         `json:"updated_at"`
	CreatedAt         *time.Time         `json:"created_at"`
	Head              *PullRequestBranch `json:"head"`
	Base              *PullRequestBranch `json:"base"`
	Mergeable         *bool              `json:"mergeable,omitempty"`
	Title             *string            `json:"title"`
	Body              *string            `json:"body"`
	User              *User              `json:"user"`
	HTMLURL           *string            `json:"html_url"`
	URL               *string            `json:"url"`
	CommitsURL        *string            `json:"commits_url"`
	ReviewCommentsURL *string            `json:"review_comments_url"`
	ReviewCommentURL  *string            `json:"review_comment_url"`
	CommentsURL       *string            `json:"comments_url"`
	StatusesURL       *string            `json:"statuses_url"`
}

func (s *PullRequestsService) List(owner, repo string) ([]PullRequest, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/pulls", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pulls := new([]PullRequest)
	resp, err := s.client.Do(req, pulls)
	if err != nil {
		return nil, resp, err
	}

	return *pulls, resp, err
}
