package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type rss struct {
	Channel channel `xml:"channel"`
}

type channel struct {
	Items []item `xml:"item"`
}

type item struct {
	Title     string `xml:"title"`
	Enclosure struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}

// TODO define somewhere else?
const rssUrl string = "https://musicforprogramming.net/rss.xml"

func GetRss() (rss, error) {
	var feed rss

	c := &http.Client{Timeout: 12 * time.Second}
	res, err := c.Get(rssUrl)
	if err != nil {
		return feed, fmt.Errorf("http call timedout: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return feed, fmt.Errorf("Failed to read body of the response: %w", err)
	}

	if err := xml.Unmarshal(body, &feed); err != nil {
		return feed, fmt.Errorf("Failed to parse XML: %w", err)
	}

	return feed, nil
}
