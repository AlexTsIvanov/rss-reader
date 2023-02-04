package main

import "net/url"

func extractValidUrls(urls []string) []string {

	var validUrls []string
	for _, val := range urls {
		_, err := url.Parse(val)
		if err != nil {
			continue
		}

		validUrls = append(validUrls, val)
	}

	return validUrls
}
