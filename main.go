package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err.Stderr))
			os.Exit(err.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "gitopen: %v", err)
		os.Exit(1)
	}
}

func run() error {
	repoURL, err := getRepositoryURL()
	if err != nil {
		return err
	}

	fmt.Printf("Opening %s\n", repoURL)

	return open(repoURL)
}

func getRepositoryURL() (string, error) {
	origin, err := getOriginURL()
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(origin, "git@") {
		return strings.TrimSuffix(origin, ".git"), nil
	}

	r := regexp.MustCompile(`^git@(.+):(.+\/.+).git$`)
	matches := r.FindStringSubmatch(origin)

	if len(matches) < 3 {
		return "", fmt.Errorf("failed to extract repository URL from %s", origin)
	}

	return fmt.Sprintf("https://%s/%s", matches[1], matches[2]), nil
}

func getOriginURL() (string, error) {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
