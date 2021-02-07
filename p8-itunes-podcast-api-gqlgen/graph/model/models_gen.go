// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type FeedItem struct {
	PubDate     string  `json:"pubDate"`
	Text        string  `json:"text"`
	Title       string  `json:"title"`
	Subtitle    string  `json:"subtitle"`
	Description string  `json:"description"`
	Image       *string `json:"image"`
	Summary     string  `json:"summary"`
	LinkURL     string  `json:"linkUrl"`
	Duration    string  `json:"duration"`
}

type Podcast struct {
	Artist       string    `json:"artist"`
	PodcastName  string    `json:"podcastName"`
	FeedURL      string    `json:"feedUrl"`
	Thumbnail    string    `json:"thumbnail"`
	EpisodeCount int       `json:"episodeCount"`
	Genres       []*string `json:"genres"`
}
