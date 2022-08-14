package main

// UrlOpener is a service that opens a URL in a default browser.
type UrlOpener struct {
	cmd Commander
}

// Open the given URL link in a browser. Actual implementation is platform-dependent.
func (opener UrlOpener) Open(link string) error {
	return opener.open(link)
}
