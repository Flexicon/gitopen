package main

import "os/exec"

type Commander interface {
	CommandOutput(name string, arg ...string) ([]byte, error)
}

type OSCommander struct{}

func (cmd OSCommander) CommandOutput(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}
