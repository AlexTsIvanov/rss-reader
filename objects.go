package main

import "github.com/AlexTsIvanov/rss-reader/pkg/rssReader"

type request struct {
	Urls []string `json:"urls"`
}

type response struct {
	Items []rssReader.RssItem `json:"items"`
}
