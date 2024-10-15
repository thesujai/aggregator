package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thesujai/aggregator/internal/database"
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

func processFeeds(db *database.Queries, concurrency int32, timeBetweenRequest time.Duration) {
	ticker := time.Tick(timeBetweenRequest)
	for ; ; <-ticker {
		feedUrlToFetch, err := db.GetNextFeedsToFetch(context.Background(), concurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var wg sync.WaitGroup
		for i := 0; i < len(feedUrlToFetch); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				feed, err := fetchFeed(feedUrlToFetch[i])
				if err != nil {
					fmt.Println(err)
				}
				for _, item := range feed.Channel.Item {
					log.Println("Found post", item.Title)
					log.Println(item.PubDate)
				}

			}()
		}
		wg.Wait()
	}
}
