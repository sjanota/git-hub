package git

import (
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

type config struct {
	repo *git.Repository
}

func (c *config) GetCurrentBranch() (string, error) {
	ref, err := c.repo.Head()
	if err != nil {
		return "", err
	}

	if !ref.Name().IsBranch() {
		return "", errors.Errorf("detached head")
	}

	return ref.Name().Short(), nil
}

func (c *config) ListRemoteNames() ([]string, error) {
	remotes, err := c.repo.Remotes()
	if err != nil {
		return nil, err
	}

	result := make([]string, len(remotes))
	for i, remote := range remotes {
		result[i] = remote.Config().Name
	}
	return result, nil
}

func (c *config) GetRemoteURL(remoteName string) (string, error) {
	remote, err := c.repo.Remote(remoteName)
	if err != nil {
		return "", err
	}

	return remote.Config().URLs[0], nil
}

func (c *config) Clean() error {
	cfg, err := c.repo.Config()
	if err != nil {
		return err
	}

	cfg.Raw.RemoveSection("github-pr")

	err = c.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}
