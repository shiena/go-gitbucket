// Copyright 2015 Mitsuhiro Koga Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

type CommitAuthor struct {
	Name  *string    `json:"name"`
	Email *string    `json:"email"`
	Date  *time.Time `json:"date"`
}

type Commit struct {
	SHA       *string       `json:"sha,omitempty"`
	Message   *string       `json:"message,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	URL       *string       `json:"url,omitempty"`
}

type RepositoryCommit struct {
	SHA       string   `json:"sha"`
	Commit    *Commit  `json:"commit"`
	Author    *User    `json:"author,omitempty"`
	Committer *User    `json:"committer,omitempty"`
	Parents   []Commit `json:"parents"`
	URL       *string  `json:"url"`
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

func (s *PullRequestsService) Get(owner, repo string, id int) (*PullRequest, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/pulls/%v", owner, repo, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pull := new(PullRequest)
	resp, err := s.client.Do(req, pull)
	if err != nil {
		return nil, resp, err
	}

	return pull, resp, err
}

func (s *PullRequestsService) ListCommits(owner, repo string, id int) ([]RepositoryCommit, *http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/pulls/%v/commits", owner, repo, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	commit := new([]RepositoryCommit)
	resp, err := s.client.Do(req, commit)
	if err != nil {
		return nil, resp, err
	}

	return *commit, resp, err
}
