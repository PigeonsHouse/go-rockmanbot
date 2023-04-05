package utils

import (
	"github.com/mattn/go-mastodon"
	"regexp"
	"strings"
)

func RemoveTagFromMastodonStatus(s *mastodon.Status) {
	tmpStr := s.Content
	rep := regexp.MustCompile(`<("[^"]*"|'[^']*'|[^'">])*>`)
	tmpStr = strings.Replace(tmpStr, "</p><p>", "\n\n", -1)
	tmpStr = strings.Replace(tmpStr, "<br />", "\n", -1)
	tmpStr = rep.ReplaceAllString(tmpStr, "")
	tmpStr = strings.Replace(tmpStr, "&apos;", "'", -1)
	tmpStr = strings.Replace(tmpStr, "&quot;", `"`, -1)
	tmpStr = strings.Replace(tmpStr, "&amp;", "&", -1)
	s.Content = tmpStr
}

func GetName(a mastodon.Account) string {
	if len(a.DisplayName) > 0 {
		return a.DisplayName
	} else {
		return a.Username
	}
}
