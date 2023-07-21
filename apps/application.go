package apps

import (
	"github.com/mattn/go-mastodon"
	"go-rockmanbot/interfaces"
	"time"
)

type Application struct {
	client    interfaces.Client
	reactions []interfaces.Reaction
	schedules []interfaces.Schedule
}

func (app *Application) Main(status *mastodon.Status) {
	for _, reaction := range app.reactions {
		if reaction.CheckTrigger(status, app.client) {
			reaction.SendReaction(status, app.client)
		}
	}
}

func (app *Application) Scheduler() {
	for range time.Tick(1 * time.Minute) {
		for _, sc := range app.schedules {
			if sc.CheckDoNow() {
				sc.Run(time.Now(), app.client)
			}
		}
	}
}

func (app *Application) SetClient(client interfaces.Client) {
	app.client = client
}
func (app *Application) SetReactions(reactions []interfaces.Reaction) {
	app.reactions = reactions
}

func (app *Application) SetSchedules(schedules []interfaces.Schedule) {
	app.schedules = schedules
}
