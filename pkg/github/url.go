package github

import (
	"fmt"
	"github.com/pkg/errors"
	std_url "net/url"
	"strings"
)

type URL struct {
	Owner          string
	RepositoryName string
	Full           string
	Path           string
}

func RepoURL(path string) (*URL, error) {
	remoteUrl := fmt.Sprintf("https://github.com/%s", path)
	return ParseURL(remoteUrl)
}

func ParseURL(s string) (*URL, error) {
	url, err := std_url.Parse(s)
	if err != nil {
		return nil, err
	}

	if url.Host != "github.com" {
		return nil, errors.Errorf("not GitHub URL %s", s)
	}

	path := strings.TrimLeft(url.Path, "/")
	pathParts := strings.Split(path, "/")
	if len(pathParts) < 2 || pathParts[0] == "" || pathParts[1] == "" {
		return nil, fmt.Errorf("invalid remote path %s", url.Path)
	}

	return &URL{
		Owner:          pathParts[0],
		RepositoryName: pathParts[1],
		Full:           s,
		Path:           path,
	}, nil
}
