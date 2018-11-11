package ops

import (
	"github.com/sjanota/git-hub/pkg/github"
	"log"
)

func Fetch() {

}

func FetchPullRequests() error {
	c := github.NewClient()
	prs, err := c.GetMyPullRequests()

	if err != nil {
		return err
	}

	for _, pr := range prs {
		log.Printf("PR-%v: %s", pr.ID, pr.Title)
	}

	return nil
}
