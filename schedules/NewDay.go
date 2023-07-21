package schedules

import (
	"github.com/mattn/go-mastodon"
	"go-rockmanbot/interfaces"
	"time"
)

type NewDay struct {
}

func (n NewDay) CheckDoNow() bool {
	scheduleTime := ScheduleTime{0, 0}
	return checkDoNow(scheduleTime)
}

func (n NewDay) Run(now time.Time, client interfaces.Client) {
	if now.Month() == 1 && now.Day() == 1 {
		// 年明けの場合(1/1)
		client.UpdateAvatar("./medias/new_icon.png")
		client.SendMessage(mastodon.Toot{
			Status: "あけましておめでとう！皆、今年もよろしくね！",
		})
	} else if now.Day() == 1 {
		// 月頭の場合
		mediaId, _ := client.SendMedia("./medias/sailormoon.jpg", false)
		client.SendMessage(mastodon.Toot{
			Status:   "月が変わってお知らせよ！",
			MediaIDs: []mastodon.ID{mastodon.ID(mediaId)},
		})
	} else {
		// 平時の場合
		client.UpdateAvatar("./medias/normal_icon.png")
		client.SendMessage(mastodon.Toot{
			Status: "日付が変わったよ！\n「＃今日やること」でトゥートすると僕が日中何度もブーストするよ！",
		})
	}
}
