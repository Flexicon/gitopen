package main

type TestCommander struct {
	Output string
	Error  error
}

func (c TestCommander) CommandOutput(name string, args ...string) ([]byte, error) {
	return []byte(c.Output), c.Error
}
