package git

import (
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

type repository struct {
	repo *git.Repository
}

func (r *repository) GetRootDir() (string, error) {
	wt, err := r.repo.Worktree()
	if err != nil {
		return "", err
	}

	return wt.Filesystem.Root(), nil
}

func (r *repository) GetCurrentBranch() (string, error) {
	ref, err := r.repo.Head()
	if err != nil {
		return "", err
	}

	if !ref.Name().IsBranch() {
		return "", errors.Errorf("detached head")
	}

	return ref.Name().Short(), nil
}

func (r *repository) ListRemoteNames() ([]string, error) {
	remotes, err := r.repo.Remotes()
	if err != nil {
		return nil, err
	}

	result := make([]string, len(remotes))
	for i, remote := range remotes {
		result[i] = remote.Config().Name
	}
	return result, nil
}

func (r *repository) GetRemoteURL(remoteName string) (string, error) {
	remote, err := r.repo.Remote(remoteName)
	if err != nil {
		return "", err
	}

	return remote.Config().URLs[0], nil
}

func (r *repository) Clean() error {
	cfg, err := r.repo.Config()
	if err != nil {
		return err
	}

	cfg.Raw.RemoveSection("github-pr")

	err = r.repo.Storer.SetConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}
