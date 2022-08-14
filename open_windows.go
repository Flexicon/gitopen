package main

import "os/exec"

func (opener UrlOpener) open(link string) error {
	return exec.Command("start", link).Run()
}
