package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Rushi2398/rssAggregator/internal/database"
)

func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scrapping on %v go routines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetFeedsNextToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error Marking feed as Fetched: ", err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error Fetching Feeds: ", err)
	}
	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post: ", item.Title, "on Feed: ", feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
