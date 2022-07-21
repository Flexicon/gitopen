package main

import "os/exec"

func open(link string) error {
	return exec.Command("open", link).Run()
}
