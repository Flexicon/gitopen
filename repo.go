package main

import (
	"fmt"
	"regexp"
	"strings"
)

type RepoService struct {
	cmd Commander
}

func (s RepoService) getRepositoryURL() (string, error) {
	origin, err := s.getOriginURL()
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

func (s RepoService) getOriginURL() (string, error) {
	out, err := s.cmd.CommandOutput("git", "remote", "get-url", "origin")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
