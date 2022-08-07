package main

import "os/exec"

func (opener UrlOpener) Open(link string) error {
	return exec.Command("start", link).Run()
}
