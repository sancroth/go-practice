package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-practice/p8-itunes-podcast-api-gqlgen/feeds"
	"go-practice/p8-itunes-podcast-api-gqlgen/graph/generated"
	"go-practice/p8-itunes-podcast-api-gqlgen/graph/model"
	"go-practice/p8-itunes-podcast-api-gqlgen/itunes"
	"go-practice/p8-itunes-podcast-api-gqlgen/utils"
)


func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.Podcast, error) {
	ias:= itunes.NewItunesApiServices()
	res,err :=ias.Search(term)
	if err!=nil{
		return nil,err
	}

	var podcasts []*model.Podcast

	for _,r := range res.Results{
		podcast:= &model.Podcast{
			Artist: r.ArtistName,
			PodcastName: r.TrackName,
			FeedURL: r.FeedURL,
			Thumbnail: r.ArtworkURL100,
			EpisodeCount: r.TrackCount,
			Genres: utils.RefStringArr(r.Genres),
		}
		podcasts = append(podcasts,podcast)
	}

	return podcasts,nil
}

func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.FeedItem, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}

	var feedItems []*model.FeedItem

	for _, item := range res.Channel.Item {
		feedItem := &model.FeedItem{
			PubDate:     item.PubDate,
			Text:        item.Text,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       utils.CheckNullString(item.Image.Href),
			Summary:     item.Summary,
			LinkURL:     item.Enclosure.URL,
			Duration:    item.Duration,
		}

		feedItems = append(feedItems, feedItem)
	}

	return feedItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }


