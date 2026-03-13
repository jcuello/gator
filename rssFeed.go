package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jcuello/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	xmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(xmlBytes, &rssFeed)
	if err != nil {
		return nil, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, rssItem := range rssFeed.Channel.Item {
		rssItem.Title = html.UnescapeString(rssItem.Title)
		rssItem.Description = html.UnescapeString(rssItem.Description)
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Feed %v is being fetched...\n", feed.Name)
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		now := time.Now()

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			pubDate, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})

		if err != nil {
			log.Println(err)
		}
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	fmt.Printf("\n%v COMPLETED!\n\n", feed.Name)
	return nil
}
