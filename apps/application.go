package apps

import (
	"github.com/mattn/go-mastodon"
	"go-rockmanbot/interfaces"
)

type Application struct {
	client    interfaces.Client
	reactions []interfaces.Reaction
}

func (app *Application) Main(status *mastodon.Status) {
	for _, reaction := range app.reactions {
		if reaction.CheckTrigger(status, app.client) {
			reaction.SendReaction(status, app.client)
		}
	}
}

func (app *Application) SetClient(client interfaces.Client) {
	app.client = client
}
func (app *Application) SetReactions(reactions []interfaces.Reaction) {
	app.reactions = reactions
}
