package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	AppName           = "gitopen"
	DefaultRemoteName = "origin"
	Version           = "v0.0.3"
)

func main() {
	cmd := OSCommander{}

	if err := run(os.Args, cmd); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			fmt.Fprint(os.Stderr, fmt.Sprintf("%s", err.Stderr))
			os.Exit(err.ExitCode())
		}

		fmt.Fprintf(os.Stderr, "%s: %v\n", AppName, err)
		os.Exit(1)
	}
}

func run(args []string, cmd Commander) error {
	arg, err := parseArg(args)
	if err != nil {
		return err
	}

	// Handle argument as flags, otherwise treat it as remote name
	switch arg {
	case "--help", "-h", "help":
		return runHelp()
	case "--version", "-v", "version":
		return runVersion()
	default:
		if strings.HasPrefix(arg, "-") {
			return fmt.Errorf("unknown option: %s", arg)
		}
	}

	repo := &RepoService{cmd}
	opener := &UrlOpener{cmd}
	return runForRemote(arg, repo, opener)
}

func runForRemote(remote string, repo *RepoService, opener *UrlOpener) error {
	repoURL, err := repo.GetRepositoryURL(remote)
	if err != nil {
		return err
	}

	fmt.Printf("Opening %s\n", repoURL)

	return opener.Open(repoURL)
}

func runHelp() error {
	helpText := fmt.Sprintf(`
%[1]s opens any Git repository remote as a link in your browser.

Usage:
  %[1]s [remote]

Flags:
  -v, --version		display the current app version
  -h, --help		display this help message
  	`, AppName)

	fmt.Println(strings.TrimSpace(helpText))
	return nil
}

func runVersion() error {
	fmt.Printf("%s %s\n", AppName, Version)
	return nil
}

func parseArg(args []string) (string, error) {
	if len(args) <= 1 {
		return DefaultRemoteName, nil
	}

	return strings.TrimSpace(args[1]), nil
}
