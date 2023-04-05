package interfaces

import (
	"github.com/mattn/go-mastodon"
)

type Reaction interface {
	CheckTrigger(s *mastodon.Status, c Client) bool
	SendReaction(s *mastodon.Status, c Client)
}
