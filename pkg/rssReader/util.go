package rssReader

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

func parseUrl(url string) ([]responseRssItem, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response rss
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Channel.Items, nil
}

func parseRssItems(responseRssItems []responseRssItem) ([]RssItem, error) {

	var rssItems []RssItem
	for _, item := range responseRssItems {
		var rssItem RssItem
		rssItem.Title = item.Title
		rssItem.Source = item.Source.Source
		rssItem.SourceUrl = item.Source.SourceURL
		rssItem.Link = item.Link
		rssItem.Description = item.Description

		// TODO - maybe handle more time formats - found multiple time formats
		if len(item.PublishDate) > 0 {
			pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PublishDate)
			if err != nil {
				return nil, err
			}
			rssItem.PublishDate = pubDate
		}

		rssItems = append(rssItems, rssItem)
	}

	return rssItems, nil
}
