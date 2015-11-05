// Copyright 2015 Mitsuhiro Koga Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitbucket

import (
	"bytes"
	"encoding/json"
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
	AutoInit      *bool   `json:"auto_init"`
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

func (s *RepositoriesService) Create(org string, repo *Repository) (*Repository, *http.Response, error) {
	var u string
	if org != "" {
		u = fmt.Sprintf("/orgs/%v/repos", org)
	} else {
		u = "/user/repos"
	}

	req, err := s.client.NewRequest("POST", u, repo)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(req, buf)
	if err != nil {
		return nil, resp, err
	}

	if buf.Len() == 0 {
		return nil, resp, err
	}

	data := buf.Bytes()
	r := new(Repository)
	json.Unmarshal(data, r)
	if r.Name != nil {
		return r, resp, err
	}

	errorResponse := &ErrorResponse{Response: resp}
	json.Unmarshal(data, errorResponse)
	return nil, resp, errorResponse
}
