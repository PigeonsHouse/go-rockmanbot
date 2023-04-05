package tests

import (
	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
	"go-rockmanbot/reactions"
	"testing"
)

var p reactions.ParrotReaction

// TestParrot_Trigger トリガーが起動するか確認するテスト
func TestParrot_Trigger(t *testing.T) {
	tc := NewTestClient()
	s := mastodon.Status{}
	s.Content = "おはようって言ってロックマン"
	assert.Equal(t, p.CheckTrigger(&s, &tc), true)
	s.Content = "おはようって言って"
	assert.Equal(t, p.CheckTrigger(&s, &tc), false)
	s.Content = "ロックマンおはようって言って"
	assert.Equal(t, p.CheckTrigger(&s, &tc), true)
	s.Content = "おはようっていってロックマン"
	assert.Equal(t, p.CheckTrigger(&s, &tc), false)
	s.Content = "ロックマンって言って"
	assert.Equal(t, p.CheckTrigger(&s, &tc), true)
}

// TestParrot_Send 返答が正しいか確認するテスト
func TestParrot_Send(t *testing.T) {
	tc := NewTestClient()
	s := mastodon.Status{}
	s.Content = "おはようって言ってロックマン"
	p.SendReaction(&s, &tc)
	assert.Equal(t, tc.GetNewestToot().Content, "おはよう")
	s.Content = "ロックマンって言って"
	p.SendReaction(&s, &tc)
	assert.Equal(t, tc.GetNewestToot().Content, "ロックマン")
	s.Content = "ロックマンおはようって言って"
	p.SendReaction(&s, &tc)
	assert.Equal(t, tc.GetNewestToot().Content, "ロックマンおはよう")
	s.Content = "って言ってロックマン"
	p.SendReaction(&s, &tc)
	assert.Equal(t, tc.GetNewestToot().Content, reactions.PARDON)
}
