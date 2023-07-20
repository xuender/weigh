package utils

import "strings"

func Has(url, key string) bool {
	if url == "" {
		return false
	}

	index := strings.Index(url, key)

	if index < 0 {
		return false
	}

	if key[0] == '/' {
		if last := index + len(key); len(url) > last {
			return url[last] == '?' || url[last] == '/'
		}
	}

	return true
}
