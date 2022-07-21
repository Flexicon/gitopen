//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package main

import "os/exec"

func open(link string) error {
	return exec.Command("xdg-open", link).Run()
}
