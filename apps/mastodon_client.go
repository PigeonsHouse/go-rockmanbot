package apps

import (
	"context"
	"github.com/mattn/go-mastodon"
)

type MastodonClient struct {
	Client *mastodon.Client
}

func (mc MastodonClient) GetTimeLine() []mastodon.Status {
	return []mastodon.Status{}
}

func (mc MastodonClient) GetMyToots() []mastodon.Status {
	return []mastodon.Status{}
}

func (mc MastodonClient) GetRanking() []mastodon.Status {
	return []mastodon.Status{}
}

func (mc MastodonClient) GetTodayTODO() []mastodon.Status {
	return []mastodon.Status{}
}

func (mc MastodonClient) GetTomorrowTODO() []mastodon.Status {
	return []mastodon.Status{}
}

func (mc MastodonClient) GetRandomFoodToot() mastodon.Status {
	return mastodon.Status{}
}

func (mc MastodonClient) SendMessage(toot mastodon.Toot) {
	mc.Client.PostStatus(context.Background(), &toot)
}

func (mc MastodonClient) RebootToot(id mastodon.ID) {
}

func (mc MastodonClient) RemoveMyMessageByID(id mastodon.ID) {
}
