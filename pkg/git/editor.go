package git

import (
	"os"
	"os/exec"
	"strings"
)

func (c *config) GetDefaultTextEditor() (string, error) {
	editor, err := getCoreEditorFromConfig()
	if err != nil {
		return "", err
	}
	if editor != "" {
		return editor, nil
	}

	editor = getEditorFromEnv()
	if editor != "" {
		return editor, nil
	}

	return getFallbackEditor(), nil

}

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
