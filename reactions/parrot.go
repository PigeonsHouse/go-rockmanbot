package reactions

import (
	"github.com/mattn/go-mastodon"
	"github.com/sirupsen/logrus"
	"go-rockmanbot/interfaces"
	"go-rockmanbot/utils"
	"strings"
)

type ParrotReaction struct {
}

const PARDON = "ごめん、もう一度言い直してくれる？"

func (p ParrotReaction) CheckTrigger(s *mastodon.Status, c interfaces.Client) bool {
	return IsCalled(s.Content) && strings.Contains(s.Content, "って言")
}

func (p ParrotReaction) SendReaction(s *mastodon.Status, c interfaces.Client) {
	logrus.Info("REACTION: オウム返し")
	newToot := s.Content
	newToot = newToot[:strings.Index(newToot, "って言")]
	newToot = strings.Replace(newToot, "@Rockmanexe", "", -1)
	newToot = strings.TrimSpace(newToot)
	logrus.Info(" [author] " + utils.GetName(s.Account))
	if len(newToot) > 0 {
		logrus.Info(" [message] " + newToot)
	} else {
		logrus.Warn(" parrot message is nothing")
		newToot = PARDON
	}
	c.SendMessage(
		mastodon.Toot{
			Status: newToot,
		},
	)
}
