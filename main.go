package main

import (
	"context"
	"github.com/mattn/go-mastodon"
	"go-rockmanbot/apps"
	"go-rockmanbot/interfaces"
	"go-rockmanbot/reactions"
	"go-rockmanbot/schedules"
	"go-rockmanbot/utils"
)

func main() {
	// .envを読み込む
	utils.LoadDotEnv()

	// ドメイン関連の初期化
	app := apps.Application{}
	mc := apps.MastodonClient{}
	mc.Client = mastodon.NewClient(&mastodon.Config{
		Server:      utils.InstanceUrl,
		AccessToken: utils.AccessToken,
	})
	app.SetClient(mc)
	app.SetReactions([]interfaces.Reaction{
		reactions.ParrotReaction{},
	})
	app.SetSchedules([]interfaces.Schedule{
		schedules.NewDay{},
	})

	// 定期実行処理
	go app.Scheduler()

	// ネイティブ関連の処理
	wsc := mc.Client.NewWSClient()
	q, _ := wsc.StreamingWSPublic(context.Background(), true)
	for e := range q {
		if t, ok := e.(*mastodon.UpdateEvent); ok {
			apps.RemoveTagFromMastodonStatus(t.Status)
			app.Main(t.Status)
		}
	}
}
