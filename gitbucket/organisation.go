package gitbucket

import (
	"net/http"
	"time"
)

type OrganisationService struct {
	client *Client
}

type Organisation struct {
	Login       *string    `json:"login"`
	Url         *string    `json:"url"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	AvatarUrl   *string    `json:"avatar_url"`
}

func (s *OrganisationService) List() ([]Organisation, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/user/orgs", nil)
	if err != nil {
		return nil, nil, err
	}

	orgs := []Organisation{}
	resp, err := s.client.Do(req, &orgs)
	if err != nil {
		return nil, resp, err
	}

	return orgs, resp, err
}

func (s *OrganisationService) ListAll() ([]Organisation, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/organizations", nil)
	if err != nil {
		return nil, nil, err
	}

	orgs := []Organisation{}
	resp, err := s.client.Do(req, &orgs)
	if err != nil {
		return nil, resp, err
	}

	return orgs, resp, err
}
