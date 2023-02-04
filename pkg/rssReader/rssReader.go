package rssReader

import (
	"fmt"
	"sync"
)

func Parse(urls []string) (results []RssItem, err error) {

	var wg sync.WaitGroup
	items := make(chan []RssItem)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			responseRssItems, err := parseUrl(url)
			if err != nil {
				fmt.Println("parseUrl err: $v", err)
				return
			}

			rssItems, err := parseRssItems(responseRssItems)
			if err != nil {
				fmt.Println("parseRssItems err: $v", err)
				return
			}

			items <- rssItems
		}(url)
	}

	go func() {
		wg.Wait()
		close(items)
	}()

	for item := range items {
		results = append(results, item...)
	}

	return results, nil

}
