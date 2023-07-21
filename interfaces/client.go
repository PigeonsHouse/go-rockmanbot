package interfaces

import "github.com/mattn/go-mastodon"

type Client interface {
	// get
	GetTimeLine() []*mastodon.Status
	GetMyToots() []*mastodon.Status
	GetRanking() []*mastodon.Status
	GetTodayTODO() []*mastodon.Status
	GetTomorrowTODO() []*mastodon.Status
	GetRandomFoodToot() *mastodon.Status
	// action
	SendMedia(filename string, isUrl bool) (string, error)
	SendMessage(toot mastodon.Toot)
	RebootToot(id mastodon.ID)
	RemoveMyMessageByID(id mastodon.ID)
	UpdateAvatar(filename string)
}
