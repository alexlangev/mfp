package episodes

import (
	"strconv"

	"github.com/alexlangev/mfp/internal/utils"
)

type episode struct {
	id    string
	title string
	url   string
}

type episodes map[int]episode

func GetEpisodes() (episodes, error) {
	eps := make(episodes)

	feed, err := utils.GetRss()
	if err != nil {
		return eps, err
	}

	rssItems := feed.Channel.Items
	for i, item := range rssItems {
		epNum := len(rssItems) - i
		eps[epNum] = episode{
			id:    strconv.Itoa(epNum),
			title: item.Title,
			url:   item.Enclosure.URL,
		}
	}

	return eps, nil
}
