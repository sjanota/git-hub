package git

import (
	"os"
	"os/exec"
	"strings"
)

func (c *config) GetDefaultTextEditor() (string, error) {
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

// Calls 'git config' because go-git does not resolve global config
func getCoreEditorFromConfig() (string, error) {
	cmd := exec.Command("git", "config", "core.editor")

	output, err := cmd.Output()
	if err != nil {
		return "", err
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
