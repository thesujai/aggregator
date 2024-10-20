package scrapper

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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

func ProcessFeeds(db *database.Queries, concurrency int32, timeBetweenRequest time.Duration) {
	ticker := time.Tick(timeBetweenRequest)
	for ; ; <-ticker {
		feedsToFetch, err := db.GetNextFeedsToFetch(context.Background(), concurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var wg sync.WaitGroup
		for _, fetchedFeed := range feedsToFetch {
			wg.Add(1)
			go func() {
				defer wg.Done()
				feed, err := fetchFeed(fetchedFeed.Url)
				if err != nil {
					fmt.Print(err)
				}
				err = db.MarkFeedFetched(context.Background(), fetchedFeed.ID)
				if err != nil {
					fmt.Println(err)
				}
				for _, item := range feed.Channel.Item {
					err := db.AddPost(context.Background(), database.AddPostParams{
						ID:          uuid.New(),
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Title:       item.Title,
						Url:         item.Link,
						Description: item.Description,
						PublishedAt: item.PubDate,
						FeedID:      fetchedFeed.ID,
					})
					if err != nil && !strings.Contains(err.Error(), "duplicate key") {
						fmt.Println("Error Occured", err)
					}
				}

			}()
		}
		wg.Wait()
	}
}
