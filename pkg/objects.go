package rssReader

import "time"

type RssItem struct {
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	SourceUrl   string    `json:"sourceUrl"`
	Link        string    `json:"link"`
	PublishDate time.Time `json:"pubDate"`
	Description string    `json:"description"`
}

type rss struct {
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Title string            `xml:"title"`
	Items []responseRssItem `xml:"item"`
}

type responseRssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishDate string `xml:"pubDate"`
	Description string `xml:"description"`
	Source      struct {
		Source    string `xml:",chardata"`
		SourceURL string `xml:"url,attr"`
	} `xml:"source"`
}
