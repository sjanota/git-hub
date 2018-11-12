package git

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

type Repo interface {
	StorePullRequest(request *PullRequest) error
	ListPullRequests() ([]*PullRequest, error)
	GetPullRequestForBranch(branch string) (*PullRequest, error)

	StaticCommentEditor(comment string, append bool) CommentEditor
	FileCommentEditor() CommentEditor

	GetRemoteURL(remote string) (string, error)
	GetRemoteForURL(url string) (string, error)
	ListRemoteNames() ([]string, error)

	GetCurrentBranch() (string, error)
	GetDefaultTextEditor() (string, error)
	GetRootDir() (string, error)
	GetCredentials(remote string) (*Credentials, error)
	Clean() error
}

func NewConfig() (Repo, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(wd)
	if err != nil {
		return nil, err
	}

	return &repository{repo}, nil
}
