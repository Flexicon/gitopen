package main

import "flag"

type Config struct {
	Remote string
	DryRun bool
}

func loadConfig() *Config {
	parsedArgs := &Config{}

	flag.StringVar(&parsedArgs.Remote, "r", DefaultRemoteName, "remote name to look up and open")
	flag.BoolVar(&parsedArgs.DryRun, "n", false, "dry run mode - don't actually open the URL just print it")
	flag.Parse()

	return parsedArgs
}
