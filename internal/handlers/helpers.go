package handlers

import (
	"os"
	"strings"
)

func EnforceHttp(url string) string {

	// if not appended with http do that
	if url[:4] != "http" {
		return "http://" + url
	}

	return url
}
func RemoveDomainError(url string) bool {
	// user trying to enter localhost

	if url == os.Getenv("DOMAIN") {
		return false
	}

	// user tries http://,https://,www,/
	// any of the above should return false

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Replace(newURL, "/", "", 1)

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
