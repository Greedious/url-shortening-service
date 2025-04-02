package helpers

import (
	"os"
	"strings"
)

var toRemovePrefixes = []string{"http://", "https://", "www."}

func EnforceHTTP(url string) string {
	if prefix := url[:4]; prefix != "http" {
		url = "http://" + url
	}
	return url
}

func RemoveDomainErrors(url string) bool {
	apiDomain := os.Getenv("API_DOMAIN")
	if url == apiDomain {
		return false
	}
	for _, prefix := range toRemovePrefixes {
		url = strings.Replace(url, prefix, "", 1)
	}
	return url != apiDomain
}
