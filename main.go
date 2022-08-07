package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	repo := &RepoService{OSCommander{}}
	opener := &UrlOpener{}

	if err := run(repo, opener); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err.Stderr))
			os.Exit(err.ExitCode())
		} else {
			fmt.Fprintf(os.Stderr, "gitopen: %v", err)
			os.Exit(1)
		}
	}
}

func run(repo *RepoService, opener *UrlOpener) error {
	repoURL, err := repo.GetRepositoryURL()
	if err != nil {
		return err
	}

	fmt.Printf("Opening %s\n", repoURL)

	return opener.Open(repoURL)
}
