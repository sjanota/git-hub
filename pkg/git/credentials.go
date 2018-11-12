package git

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"os/exec"
	"strings"
)

type Credentials struct {
	Username string
	Password string
}

func (r repository) GetCredentials(remote string) (*Credentials, error) {
	remoteURLString, err := r.GetRemoteURL(remote)
	if err != nil {
		return nil, err
	}

	remoteUrl, err := url.Parse(remoteURLString)
	if err != nil {
		return &Credentials{}, err
	}

	cmd := exec.Command("git", "credential", "fill")
	reqData := fmt.Sprintf("protocol=%s\nhost=%s\npath=%s", remoteUrl.Scheme, remoteUrl.Host, remoteUrl.Path)
	cmd.Stdin = strings.NewReader(reqData)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get user credentials: %s", output)
	}

	result := &Credentials{}
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		parts := strings.SplitN(line, "=", 2)
		if parts[0] == "username" {
			result.Username = parts[1]
		} else if parts[0] == "password" {
			result.Password = parts[1]
		}
	}

	return result, nil
}
