package main

import (
	"flag"
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
	if err := run(); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			fmt.Fprint(os.Stderr, string(err.Stderr))
			os.Exit(err.ExitCode())
		}

		fmt.Fprintf(os.Stderr, "%s: %v\n", AppName, err)
		os.Exit(1)
	}
}

func run() error {
	conf := loadConfig()
	if err := handleCommands(conf); err != nil {
		return err
	}

	cmdr := OSCommander{}
	repo := &RepoService{cmdr}
	opener := &UrlOpener{cmdr}

	return runOpenRemote(*conf, repo, opener)
}

func runOpenRemote(conf Config, repo *RepoService, opener *UrlOpener) error {
	if conf.DryRun {
		fmt.Printf("dry-run\n\n")
	}

	repoURL, err := repo.GetRepositoryURL(conf.Remote)
	if err != nil {
		return err
	}

	fmt.Printf("Opening %s\n", repoURL)

	if conf.DryRun {
		return nil
	}
	return opener.Open(repoURL)
}

func handleCommands(conf *Config) error {
	if len(os.Args) < 2 {
		return nil
	}

	arg := os.Args[1]
	if strings.HasPrefix(arg, "-") {
		return nil // commands can only be the first argument
	}

	switch arg {
	case "help":
		printHelp()
	case "version":
		printVersion()
	default:
		if conf.Remote != "" {
			return fmt.Errorf("unknown command: %s", arg)
		}
	}
	return nil
}

func printHelp() {
	helpText := fmt.Sprintf(`
%[1]s opens any Git repository remote as a link in your browser.

Commands:
  help		show this help message
  version	show version

Flags:`, AppName) + "\n"

	flag.VisitAll(func(f *flag.Flag) {
		helpText += fmt.Sprintf("  -%s=%s\t%s\n", f.Name, f.DefValue, f.Usage)
	})

	fmt.Println(strings.TrimSpace(helpText))
	os.Exit(0)
}

func printVersion() {
	fmt.Printf("%s %s\n", AppName, Version)
	os.Exit(0)
}
