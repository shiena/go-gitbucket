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
	IssuesEvent       = WebHookEvent("issues")
	IssueCommentEvent = WebHookEvent("issue_comment")
	PullRequestEvent  = WebHookEvent("pull_request")
	PushEvent         = WebHookEvent("push")
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

type WebHookPushPayload struct {
	Commits    []commit    `json:"commits"`
	Pusher     *User       `json:"pusher"`
	Ref        *string     `json:"ref"`
	Repository *Repository `json:"repository"`
}
