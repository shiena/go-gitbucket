package gitbucket

import (
	"fmt"
	"net/http"
	"time"
)

type IssueComment struct {
	ID        *int       `json:"id,omitempty"`
	User      *User      `json:"user,omitempty"`
	Body      *string    `json:"body"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	HTMLURL   *string    `json:"html_url,omitempty"`
}

type IssuesService struct {
	client *Client
}

func (s *IssuesService) ListComments(owner, repo string, id int) ([]IssueComment, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/issues/%v/comments", owner, repo, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	comments := new([]IssueComment)
	resp, err := s.client.Do(req, comments)
	if err != nil {
		return nil, resp, err
	}

	return *comments, resp, err
}

func (s *IssuesService) CreateComment(owner, repo string, id int, comment *IssueComment) (*IssueComment, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/issues/%v/comments", owner, repo, id)
	req, err := s.client.NewRequest("POST", u, comment)
	if err != nil {
		return nil, nil, err
	}

	c := new(IssueComment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
