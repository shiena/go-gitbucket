// Copyright 2015 Mitsuhiro Koga Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitbucket

import (
	"fmt"
	"net/http"
	"time"
)

type UsersService struct {
	client *Client
}

// User represents a API user.
type User struct {
	Login     *string    `json:"login"`
	Email     *string    `json:"email"`
	Type      *string    `json:"type"`
	SiteAdmin *bool      `json:"site_admin"`
	CreatedAt *time.Time `json:"created_at"`
	URL       *string    `json:"url"`
	HTMLURL   *string    `json:"html_url"`
}

func (s *UsersService) Get(user string) (*User, *http.Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("/users/%v", user)
	} else {
		u = "/user"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

func (s *UsersService) GetAll() ([]User, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/users", nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := []User{}
	resp, err := s.client.Do(req, &uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

