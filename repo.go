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
	remoteURL, err := s.getRemoteURL(remote)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(remoteURL, "http") && s.validURL(remoteURL) {
		return s.handleHttpRemote(remoteURL)
	}

	if strings.HasPrefix(remoteURL, "ssh://git@") {
		return s.handleSSHRemote(remoteURL)
	}

	return s.handleGitURLRemote(remoteURL)
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

func (s RepoService) handleHttpRemote(remoteURL string) (string, error) {
	return strings.TrimSuffix(remoteURL, ".git"), nil
}

func (s RepoService) handleSSHRemote(remoteURL string) (string, error) {
	parts := strings.Split(remoteURL, "ssh://git@")

	if len(parts) < 2 {
		return "", fmt.Errorf("failed to extract repository URL from %s", remoteURL)
	}
	return strings.TrimSuffix(fmt.Sprintf("https://%s", parts[1]), ".git"), nil
}

func (s RepoService) handleGitURLRemote(remoteURL string) (string, error) {
	r := regexp.MustCompile(`^git@(.+):(.+\/.+).git$`)
	matches := r.FindStringSubmatch(remoteURL)

	if len(matches) < 3 {
		return "", fmt.Errorf("failed to extract repository URL from %s", remoteURL)
	}
	return fmt.Sprintf("https://%s/%s", matches[1], matches[2]), nil
}
