package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AlexTsIvanov/rss-reader/pkg/rssReader"
	"github.com/labstack/echo/v4"
)

type RssService struct {
	*echo.Echo
}

func NewRssService() *RssService {
	e := echo.New()
	return &RssService{e}
}

func (s RssService) handleRssFeed(c echo.Context) error {
	// read request body
	if c.Request().Body == nil {
		return fmt.Errorf("nil body")
	}

	defer c.Request().Body.Close()
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		body = make([]byte, 0)
	}

	var req request
	err = json.Unmarshal(body, &req)
	if err != nil {
		return fmt.Errorf("couldn't parse request body, error: %v, raw body: %s",
			err, string(body))
	}

	validUrls := extractValidUrls(req.Urls)

	items, err := rssReader.Parse(validUrls)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("couldnt parse rss feed")
	}

	return c.JSON(http.StatusOK, items)
}
