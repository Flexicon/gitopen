package main

type UrlOpener struct{}

func (opener UrlOpener) Open(url string) error {
	return open(url)
}
