package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/minzhoudu/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func convertDbUserToUserDto(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func convertToFeedDto(dbFeed database.Feed) Feed {
	// return Feed{
	// 	ID:        dbFeed.ID,
	// 	CreatedAt: dbFeed.CreatedAt,
	// 	UpdatedAt: dbFeed.UpdatedAt,
	// 	Name:      dbFeed.Name,
	// 	Url:       dbFeed.Url,
	// 	UserID:    dbFeed.UserID,
	// }
	return Feed(dbFeed)
}

func convertToFeedsSliceDto(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, Feed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func convertToFeedFollowsSliceDto(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, FeedFollow(dbFeedFollow))
	}

	return feedFollows
}