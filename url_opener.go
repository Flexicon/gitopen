package main

// UrlOpener is a service that opens a URL in the default browser.
//
// The `Open(url string)` method implementation is platform-dependent.
type UrlOpener struct {
	cmd Commander
}
