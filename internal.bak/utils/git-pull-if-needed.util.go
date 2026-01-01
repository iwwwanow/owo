package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GitPullIfNeeded(dirPath string) error {
	fmt.Printf("Attempting git pull from: %s\n", dirPath)

	// Получаем текущий remote URL
	remoteCmd := exec.Command("git", "-C", dirPath, "remote", "get-url", "origin")
	remoteOutput, err := remoteCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get remote URL: %v", err)
	}

	remoteURL := strings.TrimSpace(string(remoteOutput))

	// Если это SSH URL, меняем на HTTPS
	if strings.Contains(remoteURL, "git@github.com") {
		httpsURL := strings.Replace(remoteURL, "git@github.com:", "https://github.com/", 1)
		setCmd := exec.Command("git", "-C", dirPath, "remote", "set-url", "origin", httpsURL)
		if err := setCmd.Run(); err != nil {
			return fmt.Errorf("failed to set HTTPS URL: %v", err)
		}
		fmt.Printf("Changed remote URL to HTTPS: %s\n", httpsURL)
	}

	// Выполняем pull
	cmd := exec.Command("git", "-C", dirPath, "pull", "--rebase")
	output, err := cmd.CombinedOutput()
	fmt.Printf("Git pull output: %s\n", string(output))

	if err != nil {
		return fmt.Errorf("git pull failed: %v", err)
	}

	return nil
}
