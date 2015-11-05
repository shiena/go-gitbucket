[![GoDoc](https://godoc.org/github.com/shiena/go-gitbucket/gitbucket?status.svg)](https://godoc.org/github.com/shiena/go-gitbucket/gitbucket)

# go-gitbucket

go-gitbucket is a Go client library for accessing the [GitBucket API](https://github.com/takezoe/gitbucket/wiki/API-WebHook).

## Usage

```go
import "github.com/shiena/go-gitbucket/gitbucket"
```

GitBucket API

```go
url := "http://privatesite.example.org:8080/gitbucket/"
accessToken := ... // Generated access token from your account settings

client := gitbucket.NewClient(url, accessToken)

// Users API
user, res, err := client.UsersService.Get("username")

// Repositories API
repo, res, err := client.RepositoriesService.Get("username", "reponame")
repos, res, err := client.RepositoriesService.ListStatuses("username", "master")
stats, res, err := client.RepositoriesService.GetCombinedStatus("username", "master")
stat, res, err := client.RepositoriesService.CreateStatus(
	"username",
	"master",
	&gitbucket.RepoStatus{
		State: gitbucket.String("success"), // Required
		TargetURL: gitbucket.String("https://examples.org/build/status"),
		Description: gitbucket.String("The build succeeded"),
		Context: gitbucket.String("continuous-integration/jenkins"),
	},
)

// Issues API
issues, res, err := client.IssuesService.ListComments("username", "reponame", 1)
issue, res, err := client.IssuesService.CreateComment(
	"username",
	"reponame",
	1,
	&gitbucket.IssueComment{
		Body: gitbucket.String("Me too"),
	},
)

// Pull Requests API
pull, res, err := client.PullRequestService.Get("username", "reponame", 1)
pulls, res, err := client.PullRequestService.List("username", "reponame")
commits, res, err := client.PullRequestService.ListCommits("username", "reponame", 1)

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
	"github.com/shiena/go-gitbucket/gitbucket"
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

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.

