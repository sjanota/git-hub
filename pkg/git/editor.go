package git

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
)

func (r *repository) GetDefaultTextEditor() (string, error) {
	if editor, err := getCoreEditorFromConfig(); err != nil {
		return "", err
	} else if editor != "" {
		return editor, nil
	}

	if editor := getEditorFromEnv(); editor != "" {
		return editor, nil
	}

	return getFallbackEditor(), nil
}

// Calls 'git repository' because go-git does not resolve global repository
func getCoreEditorFromConfig() (string, error) {
	cmd := exec.Command("git", "config", "core.editor")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", errors.Wrapf(err, "cannot read git config: %s", output)
	}

	return strings.TrimSpace(string(output)), nil
}

func getEditorFromEnv() string {
	return os.Getenv("EDITOR")
}

func getFallbackEditor() string {
	_, err := exec.LookPath("vim")
	if err != nil {
		return "vi"
	}
	return "vim"
}
