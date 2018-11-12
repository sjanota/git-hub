package git

import "regexp"

var (
	urlGitSuffixPattern = regexp.MustCompile(`\.git$`)
)

type RemotesLister interface {
	List(cfg Repo) ([]string, error)
}

type AllRemotesLister struct{}

func (AllRemotesLister) List(repo Repo) ([]string, error) {
	return repo.ListRemoteNames()
}

type OneRemoteLister struct {
	Remote string
}

func (l OneRemoteLister) List(repo Repo) ([]string, error) {
	return []string{l.Remote}, nil
}

func (r *repository) GetRemoteURL(remoteName string) (string, error) {
	remote, err := r.repo.Remote(remoteName)
	if err != nil {
		return "", err
	}

	url := remote.Config().URLs[0]
	url = string(urlGitSuffixPattern.ReplaceAll([]byte(url), []byte{}))

	return url, nil
}

func (r *repository) GetRemoteForURL(url string) (string, error) {
	remotes, err := r.repo.Remotes()
	if err != nil {
		return "", err
	}

	for _, remote := range remotes {
		if remote.Config().URLs[0] == url || remote.Config().URLs[0] == (url+".git") {
			return remote.Config().Name, nil
		}
	}

	return "", NoRemoteForURL{Url: url}
}
