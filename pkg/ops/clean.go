package ops

import "github.com/sjanota/git-hub/pkg/config"

func Clean() error {
	cfg, err := config.NewGitConfig()
	if err != nil {
		return err
	}

	return cfg.Clean()
}
