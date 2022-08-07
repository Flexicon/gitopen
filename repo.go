package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type RepoService struct {
	cmd Commander
}

func (s RepoService) GetRepositoryURL(remote string) (string, error) {
	origin, err := s.getRemoteURL(remote)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(origin, "git@") && s.validURL(origin) {
		return strings.TrimSuffix(origin, ".git"), nil
	}

	r := regexp.MustCompile(`^git@(.+):(.+\/.+).git$`)
	matches := r.FindStringSubmatch(origin)

	if len(matches) < 3 {
		return "", fmt.Errorf("failed to extract repository URL from %s", origin)
	}

	return fmt.Sprintf("https://%s/%s", matches[1], matches[2]), nil
}

func (s RepoService) getRemoteURL(remote string) (string, error) {
	out, err := s.cmd.CommandOutput("git", "remote", "get-url", remote)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func (s RepoService) validURL(input string) bool {
	_, err := url.ParseRequestURI(input)

	return err == nil
}
