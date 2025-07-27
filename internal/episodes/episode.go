package episodes

import (
	"strconv"

	"github.com/alexlangev/mfp/internal/utils"
)

type episode struct {
	Id    string
	Title string
	Url   string
}

type Episodes []episode

func GetEpisodes() (Episodes, error) {
	feed, err := utils.GetRss()
	if err != nil {
		return nil, err

	}

	rssItems := feed.Channel.Items
	eps := make([]episode, len(rssItems))

	for i, item := range rssItems {
		epNum := len(rssItems) - i

		eps[epNum-1] = episode{
			Id:    strconv.Itoa(epNum),
			Title: item.Title,
			Url:   item.Enclosure.URL,
		}
	}

	return eps, nil
}
