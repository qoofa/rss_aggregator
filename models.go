package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/qoofa/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserTOUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
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

func databaseFeedToFeed(f database.Feed) Feed {
	return Feed{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Url:       f.Url,
		UserID:    f.UserID,
	}
}

func databaseFeedsToFeeds(f []database.Feed) []Feed {
	feeds := []Feed{}
	for _, v := range f {
		feeds = append(feeds, databaseFeedToFeed(v))
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

func databaseFeedFollowToFeedFollow(f database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		UserID:    f.UserID,
		FeedID:    f.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(f []database.FeedFollow) []FeedFollow {
	d := []FeedFollow{}
	for _, v := range f {
		d = append(d, databaseFeedFollowToFeedFollow(v))
	}
	return d
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(d database.Post) Post {
	var description *string
	if d.Description.Valid {
		description = &d.Description.String
	}

	return Post{
		ID:          d.ID,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		Title:       d.Title,
		Description: description,
		PublishedAt: d.PublishedAt,
		Url:         d.Url,
		FeedID:      d.FeedID,
	}
}

func databasePostsToPosts(d []database.Post) []Post {
	posts := []Post{}
	for _, v := range d {
		posts = append(posts, databasePostToPost(v))
	}
	return posts
}
