package main

import (
	"context"
	"github.com/mattn/go-mastodon"
	"go-rockmanbot/apps"
	"go-rockmanbot/interfaces"
	"go-rockmanbot/reactions"
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

	// ネイティブ関連の処理
	wsc := mc.Client.NewWSClient()
	q, _ := wsc.StreamingWSPublic(context.Background(), true)
	for e := range q {
		if t, ok := e.(*mastodon.UpdateEvent); ok {
			utils.RemoveTagFromMastodonStatus(t.Status)
			app.Main(t.Status)
		}
	}
}
