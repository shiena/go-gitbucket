// Copyright 2015 Mitsuhiro Koga Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitbucket

import (
	"net/http"
	"time"
)

type WebHookEvent string

const (
	IssuesEvent                   = WebHookEvent("issues")
	IssueCommentEvent             = WebHookEvent("issue_comment")
	PullRequestEvent              = WebHookEvent("pull_request")
	PullRequestReviewCommentEvent = WebHookEvent("pull_request_review_comment")
	PushEvent                     = WebHookEvent("push")
)

func GetWebHookEvent(req *http.Request) WebHookEvent {
	return WebHookEvent(req.Header.Get("X-Github-Event"))
}

type commit struct {
	ID        *string       `json"id"`
	Message   *string       `json:"message"`
	Timestamp *time.Time    `json:"timestamp"`
	Added     []string      `json:"added"`
	Removed   []string      `json:"removed"`
	Modified  []string      `json:"modified"`
	Author    *CommitAuthor `json:"author"`
	Committer *CommitAuthor `json:"committer"`
	URL       *string       `json:"url"`
}

type issue struct {
	Number      *int       `json:"number"`
	Title       *string    `json:"title"`
	User        *User      `json:"user"`
	State       *string    `json:"state"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Body        *string    `json:"body"`
	CommentsURL *string    `json:"comments_url"`
	HTMLURL     *string    `json:"html_url"`
}

type path struct {
	Href *string `json:"href"`
}

type links struct {
	Self        *path `json:"self"`
	HTML        *path `json:"html"`
	PullRequest *path `json:"pull_request"`
}

type comment struct {
	ID             *int       `json:"id"`
	Path           *string    `json:"path"`
	CommitID       *string    `json:"commit_id"`
	User           *User      `json:"user"`
	Body           *string    `json:"body"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	URL            *string    `json:"url"`
	HTMLURL        *string    `json:"html_url"`
	PullRequestURL *string    `json:"pull_request_url"`
	Links          *links     `json:"_links"`
}

// Action is one of "opened", "reopen", "closed"
type WebHookIssuesPayload struct {
	Action     *string     `json:"action"`
	Sender     *User       `json:"sender"`
	Number     *int        `json:"number"`
	Issue      *issue      `json:"issue"`
	Repository *Repository `json:"repository"`
}

// Action is only "created"
type WebHookIssueCommentPayload struct {
	Action     *string       `json:"action"`
	Sender     *User         `json:"sender"`
	Issue      *issue        `json:"issue"`
	Comment    *IssueComment `json:"comment"`
	Repository *Repository   `json:"repository"`
}

// Action is one of "created", "synchronize"
type WebHookPullRequestPayload struct {
	Action      *string      `json:"action"`
	Sender      *User        `json:"sender"`
	Number      *int         `json:"number"`
	PullRequest *PullRequest `json:"pull_request"`
	Repository  *Repository  `json:"repository"`
}

// Action is only "created"
type WebHookPullRequestReviewCommentPayload struct {
	Action      *string      `json:"action"`
	Comment     *comment     `json:"comment"`
	PullRequest *PullRequest `json:"pull_request"`
	Repository  *Repository  `json:"repository"`
	Sender      *User        `json:"sender"`
}

type WebHookPushPayload struct {
	Commits    []commit    `json:"commits"`
	Pusher     *User       `json:"pusher"`
	Ref        *string     `json:"ref"`
	Before     *string     `json:"before"`
	After      *string     `json:"after"`
	Repository *Repository `json:"repository"`
	Compare    *string     `json:"compare"`
	HeadCommit *commit     `json:"head_commit"`
}
