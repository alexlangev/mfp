package episodes

import (
	"strconv"

	"github.com/alexlangev/mfp/internal/utils"
)

type Episode struct {
	Id    string
	Title string
	Url   string
}

type Episodes []Episode

func GetEpisodes() (Episodes, error) {
	feed, err := utils.GetRss()
	if err != nil {
		return nil, err

	}

	rssItems := feed.Channel.Items
	eps := make([]Episode, len(rssItems))

	for i, item := range rssItems {
		epNum := len(rssItems) - i

		eps[epNum-1] = Episode{
			Id:    strconv.Itoa(epNum),
			Title: item.Title,
			Url:   item.Enclosure.URL,
		}
	}

	return eps, nil
}
