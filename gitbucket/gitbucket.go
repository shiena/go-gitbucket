package gitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-gitbucket/" + libraryVersion
	mediaTypeJSON  = "application/json"
)

func Version() string {
	return libraryVersion
}

// Client represents a GitBucket API client.
type Client struct {
	accessToken  string
	client       *http.Client
	BaseURL      *url.URL
	UserAgent    string
	Users        *UsersService
	Repositories *RepositoriesService
	Pullrequests *interface{}
	Issues       *interface{}
}

// NewClient initializes and returns a API client.
func NewClient(urlStr, token string) (*Client, error) {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	client := &Client{
		BaseURL:     baseURL,
		accessToken: token,
		client:      &http.Client{},
		UserAgent:   userAgent,
	}
	client.Users = &UsersService{client}
	client.Repositories = &RepositoriesService{client}
	return client, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse("api/v3" + urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaTypeJSON)
	if buf != nil {
		req.Header.Add("Content-Type", mediaTypeJSON)
	}
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if c.accessToken != "" {
		req.Header.Set("Authorization", "token "+c.accessToken)
	}
	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}

type ErrorResponse struct {
	Response         *http.Response // HTTP response that caused this error
	Message          string         `json:"message"` // error message
	DocumentationURL *string        `json:"documentation_url,omitempty"`
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	if strings.HasPrefix(r.Header.Get("Content-Type"), mediaTypeJSON) {
		data, err := ioutil.ReadAll(r.Body)
		if err == nil && data != nil {
			json.Unmarshal(data, errorResponse)
		}
	}
	return errorResponse
}

type RateLimit struct {
}

func (c *Client) RateLimit() (*RateLimit, *http.Response, error) {
	req, err := c.NewRequest("GET", "/rate_limit", nil)
	if err != nil {
		return nil, nil, err
	}

	rResp := new(RateLimit)
	resp, err := c.Do(req, rResp)
	if err != nil {
		return nil, resp, err
	}

	return rResp, resp, err
}
