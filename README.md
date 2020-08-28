# go-gitbucket

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shiena/go-gitbucket/gitbucket)](https://pkg.go.dev/github.com/shiena/go-gitbucket/gitbucket)

go-gitbucket is a Go client library for accessing the [GitBucket API](https://github.com/takezoe/gitbucket/wiki/API-WebHook).

## Usage

```go
import "github.com/StingrayDigital/go-gitbucket/gitbucket"
```

GitBucket API

```go
url := "http://privatesite.example.org:8080/gitbucket/"
accessToken := "... your access token ..." // Generated from your account settings

client, err := gitbucket.NewClient(url, accessToken)

// Users API
user, res, err := client.Users.Get("username")

// Repositories API
repo, res, err := client.Repositories.Get("username", "reponame", "ref")
repos, res, err := client.Repositories.ListStatuses("username", "master", "ref")
stats, res, err := client.Repositories.GetCombinedStatus("username", "master", "ref")
stat, res, err := client.Repositories.CreateStatus(
	"username",
	"master",
	"sha",
	&gitbucket.RepoStatus{
		State: gitbucket.String("success"), // Required
		TargetURL: gitbucket.String("https://examples.org/build/status"),
		Description: gitbucket.String("The build succeeded"),
		Context: gitbucket.String("continuous-integration/jenkins"),
	},
)
newRepo, res, err := client.Repositories.Create(
	"", // if parameter is not empty, set group name.
	&gitbucket.Repository{
		Name: gitbucket.String("Hello-World"), // Required
		Description: gitbucket.String("This is your first repository"),
		Private: gitbucket.Bool(true), // Default: false
		AutoInit: gitbucket.Bool(true), // Default: false
	},
)

// Issues API
issues, res, err := client.Issues.ListComments("username", "reponame", 1)
issue, res, err := client.Issues.CreateComment(
	"username",
	"reponame",
	1,
	&gitbucket.IssueComment{
		Body: gitbucket.String("Me too"),
	},
)

// Pull Requests API
pull, res, err := client.PullRequests.Get("username", "reponame", 1)
pulls, res, err := client.PullRequests.List("username", "reponame")
commits, res, err := client.PullRequests.ListCommits("username", "reponame", 1)

// Others API
rateLimit, res, err := client.RateLimit()
```

Receive a GitBucket WebHook Events

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/StingrayDigital/go-gitbucket/gitbucket"
)

func getEvent(r *http.Request) interface{} {
	event := gitbucket.GetWebHookEvent(r)
	switch event {
	case gitbucket.IssuesEvent:
		return &gitbucket.WebHookIssuesPayload{}
	case gitbucket.IssueCommentEvent:
		return &gitbucket.WebHookIssueCommentPayload{}
	case gitbucket.PullRequestEvent:
		return &gitbucket.WebHookPullRequestPayload{}
	case gitbucket.PullRequestReviewCommentEvent:
		return &gitbucket.WebHookPullRequestReviewCommentPayload{}
	case gitbucket.PushEvent:
		return &gitbucket.WebHookPushPayload{}
	default:
		return nil
	}
}

func printEventHandler(w http.ResponseWriter, r *http.Request) {
	payload := r.FormValue("payload")
	event := getEvent(r)
	if payload != "" && event != nil {
		if err := json.Unmarshal([]byte(payload), event); err == nil {
			log.Println(pp.Sprint(event))
		}
	} else {
		log.Println("Received a invalid message")
	}
}

func main() {
	http.HandleFunc("/", printEventHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

## Credits

- @toVersus
    - Fix sample code
- @StingrayDigital
    - Added organisations listing
    - Added users listing
    - Added go module support
- @cooperspencer
    - Added get user repo
	- Added sshurl

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.

