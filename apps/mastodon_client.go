package apps

import (
	"context"
	"encoding/base64"
	"github.com/mattn/go-mastodon"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type MastodonClient struct {
	Client *mastodon.Client
}

func (mc MastodonClient) GetTimeLine() []*mastodon.Status {
	return []*mastodon.Status{}
}

func (mc MastodonClient) GetMyToots() (myToots []*mastodon.Status) {
	myToots, _ = mc.Client.GetAccountStatuses(context.Background(), "4", nil)
	return
}

func (mc MastodonClient) GetRanking() []*mastodon.Status {
	return []*mastodon.Status{}
}

func (mc MastodonClient) GetTodayTODO() []*mastodon.Status {
	return []*mastodon.Status{}
}

func (mc MastodonClient) GetTomorrowTODO() []*mastodon.Status {
	return []*mastodon.Status{}
}

func (mc MastodonClient) GetRandomFoodToot() *mastodon.Status {
	return &mastodon.Status{}
}

func (mc MastodonClient) SendMedia(filename string, isUrl bool) (string, error) {
	fn := ""
	if isUrl {
		res, err := http.Get(filename)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		fn = "temp" + filepath.Ext(filename)
		f, err := os.Create(fn)
		_, err = io.Copy(f, res.Body)
	} else {
		fn = filename
	}
	media, err := mc.Client.UploadMedia(context.Background(), fn)
	if err != nil {
		return "", err
	}
	err = os.Remove(fn)
	if err != nil {
		return "", err
	}
	return string(media.ID), nil
}

func (mc MastodonClient) SendMessage(toot mastodon.Toot) {
	mc.Client.PostStatus(context.Background(), &toot)
}

func (mc MastodonClient) RebootToot(id mastodon.ID) {
	mc.Client.Reblog(context.Background(), id)
}

func (mc MastodonClient) RemoveMyMessageByID(id mastodon.ID) {
	mc.Client.DeleteStatus(context.Background(), id)
}

func (mc MastodonClient) UpdateAvatar(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return
	}
	data := make([]byte, fi.Size())
	_, err = file.Read(data)
	if err != nil {
		return
	}
	encodedStr := base64.StdEncoding.EncodeToString(data)

	mc.Client.AccountUpdate(context.Background(), &mastodon.Profile{
		Avatar: "data:image/png;base64," + encodedStr,
	})
}

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
