package internal

import (
	"os"
	"os/exec"
)


func EditInEditor(path string, initialText string) (string, error) {
	// Get editor from env, fallback to vi
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	// Write initial text to the given path
	if err := os.WriteFile(path, []byte(initialText), 0644); err != nil {
		return "", err
	}

	// Open editor
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Read back edited file
	edited, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(edited), nil
}
