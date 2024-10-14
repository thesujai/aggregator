package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println("Got an error while fetching feed: ", err)
		return RSSFeed{}, err

	}
	defer resp.Body.Close()
	feed := RSSFeed{}
	err = xml.NewDecoder(resp.Body).Decode(&feed)
	if err != nil {
		fmt.Println("Error while decoding xml", err)
		return RSSFeed{}, err
	}
	return feed, nil
}
