package tests

import (
	"github.com/mattn/go-mastodon"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
	"time"
)

type TestClient struct {
	timeLine    []*mastodon.Status
	mediaList   []mastodon.Attachment
	accountList []mastodon.Account
	cookingTag  mastodon.Tag
}

func NewTestClient() TestClient {
	tc := TestClient{}
	tc.accountList = []mastodon.Account{
		{
			ID:          "1",
			Username:    "Rockmanexe",
			DisplayName: "ロックマンエグゼ",
			CreatedAt:   time.Date(2020, 8, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "2",
			Username:    "Pigeon_house2020",
			DisplayName: "鳩屋敷",
			CreatedAt:   time.Date(2020, 5, 10, 0, 0, 0, 0, time.UTC),
		},
	}
	rockman := tc.accountList[0]
	hato := tc.accountList[1]
	tc.mediaList = []mastodon.Attachment{{
		ID:   "1",
		Type: "image",
		URL:  "https://s3.ap-northeast-1.wasabisys.com/mastodondb/media_attachments/files/109/308/194/603/535/371/original/9363cefc6a6b45b9.jpeg",
		Meta: mastodon.AttachmentMeta{
			Original: mastodon.AttachmentSize{
				Width:  1247,
				Height: 1663,
				Size:   "1247x1663",
				Aspect: 0.7498496692723993,
			},
		},
	}}
	curry := []mastodon.Attachment{tc.mediaList[0]}
	tc.cookingTag = mastodon.Tag{
		Name:    "CompositeCookingClub",
		URL:     "https://mastodon.compositecomputer.club/tags/CompositeCookingClub",
		History: []mastodon.History{},
	}
	tc.timeLine = []*mastodon.Status{
		{
			ID:        "1",
			Account:   rockman,
			Content:   "どうも",
			CreatedAt: time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
		},
		{
			ID:               "2",
			Account:          hato,
			Content:          `<p>皆既月食<br /><a href="https://mastodon.compositecomputer.club/tags/CompositeCookingClub" cl"mention hashtag" rel="tag">#<span>CompositeCookingClub</span></a></p>`,
			CreatedAt:        time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
			MediaAttachments: curry,
			Tags:             []mastodon.Tag{tc.cookingTag},
		},
		{
			ID:        "3",
			Account:   rockman,
			Content:   "どうも",
			CreatedAt: time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
		},
		{
			ID:               "4",
			Account:          hato,
			Content:          `<p>皆既月食<br /><a href="https://mastodon.compositecomputer.club/tags/CompositeCookingClub" cl"mention hashtag" rel="tag">#<span>CompositeCookingClub</span></a></p>`,
			CreatedAt:        time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
			MediaAttachments: curry,
			Tags:             []mastodon.Tag{tc.cookingTag},
		},
		{
			ID:        "5",
			Account:   rockman,
			Content:   "どうも",
			CreatedAt: time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
		},
		{
			ID:               "6",
			Account:          hato,
			Content:          `<p>皆既月食<br /><a href="https://mastodon.compositecomputer.club/tags/CompositeCookingClub" cl"mention hashtag" rel="tag">#<span>CompositeCookingClub</span></a></p>`,
			CreatedAt:        time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
			MediaAttachments: curry,
			Tags:             []mastodon.Tag{tc.cookingTag},
		},
		{
			ID:        "7",
			Account:   rockman,
			Content:   "どうも",
			CreatedAt: time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
		},
		{
			ID:               "8",
			Account:          hato,
			Content:          `<p>皆既月食<br /><a href="https://mastodon.compositecomputer.club/tags/CompositeCookingClub" cl"mention hashtag" rel="tag">#<span>CompositeCookingClub</span></a></p>`,
			CreatedAt:        time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
			MediaAttachments: curry,
			Tags:             []mastodon.Tag{tc.cookingTag},
		},
		{
			ID:               "9",
			Account:          hato,
			Content:          `<p>皆既月食<br /><a href="https://mastodon.compositecomputer.club/tags/CompositeCookingClub" cl"mention hashtag" rel="tag">#<span>CompositeCookingClub</span></a></p>`,
			CreatedAt:        time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
			MediaAttachments: curry,
			Tags:             []mastodon.Tag{tc.cookingTag},
		},
		{
			ID:               "10",
			Account:          hato,
			Content:          `<p>皆既月食<br /><a href="https://mastodon.compositecomputer.club/tags/CompositeCookingClub" cl"mention hashtag" rel="tag">#<span>CompositeCookingClub</span></a></p>`,
			CreatedAt:        time.Date(2022, 11, 8, 12, 31, 26, 975000000, time.UTC),
			MediaAttachments: curry,
			Tags:             []mastodon.Tag{tc.cookingTag},
		},
	}
	return tc
}

func (tc *TestClient) GetTimeLine() []*mastodon.Status {
	return tc.timeLine
}

func (tc *TestClient) GetNewestToot() *mastodon.Status {
	return tc.timeLine[len(tc.timeLine)-1]
}

func (tc *TestClient) GetMyToots() (myStatuses []*mastodon.Status) {
	for _, t := range tc.timeLine {
		if t.Account.Username == "Rockmanexe" {
			myStatuses = append(myStatuses, t)
		}
	}
	return
}

func (tc *TestClient) GetRanking() (statuses []*mastodon.Status) {
	statuses = append([]*mastodon.Status{}, tc.timeLine...)
	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].RepliesCount*5+statuses[i].ReblogsCount*2+statuses[i].FavouritesCount >
			statuses[j].RepliesCount*5+statuses[j].ReblogsCount*2+statuses[j].FavouritesCount
	})
	statuses = statuses[:5]
	return
}

func (tc *TestClient) GetTodayTODO() (statuses []*mastodon.Status) {
	for _, t := range tc.timeLine {
		for _, tg := range t.Tags {
			if tg.Name == "明日やること" {
				statuses = append(statuses, t)
				break
			}
		}
	}
	return
}

func (tc *TestClient) GetTomorrowTODO() (statuses []*mastodon.Status) {
	for _, t := range tc.timeLine {
		for _, tg := range t.Tags {
			if tg.Name == "明日やること" {
				statuses = append(statuses, t)
				break
			}
		}
	}
	return
}

func (tc *TestClient) GetRandomFoodToot() *mastodon.Status {
	for _, t := range tc.timeLine {
		if len(t.MediaAttachments) > 0 {
			for _, tg := range t.Tags {
				if tg.Name == "CompositeCookingClub" {
					return t
				}
			}
		}
	}
	return &mastodon.Status{}
}

func (tc *TestClient) SendMedia(filename string, isUrl bool) (string, error) {
	url := ""
	if isUrl {
		url = "https://s3.ap-northeast-1.wasabisys.com/mastodondb/media_attachments/files/109/308/194/603/535/371/original/9363cefc6a6b45b9.jpeg"
	} else {
		url = filename
	}
	newId := strconv.Itoa(len(tc.mediaList) + 1)
	newMedia := mastodon.Attachment{
		ID:   mastodon.ID(newId),
		Type: "image",
		URL:  url,
		Meta: mastodon.AttachmentMeta{
			Original: mastodon.AttachmentSize{
				Width:  1247,
				Height: 1663,
				Size:   "1247x1663",
				Aspect: 0.7498496692723993,
			},
		},
	}
	tc.mediaList = append(tc.mediaList, newMedia)
	return newId, nil
}

func (tc *TestClient) SendMessage(toot mastodon.Toot) {
	maxId := tc.timeLine[len(tc.timeLine)-1].ID
	numId, _ := strconv.Atoi(string(maxId))

	var attachment []mastodon.Attachment
	for _, media := range tc.mediaList {
		if slices.Contains(toot.MediaIDs, media.ID) {
			attachment = append(attachment, media)
		}
	}

	tc.timeLine = append(tc.timeLine, &mastodon.Status{
		ID:               mastodon.ID(strconv.Itoa(numId + 1)),
		CreatedAt:        time.Now(),
		Content:          toot.Status,
		MediaAttachments: attachment,
	})
}

func (tc *TestClient) RebootToot(id mastodon.ID) {
	var targetToot mastodon.Status
	for _, t := range tc.timeLine {
		if t.ID == id {
			targetToot = *t
		}
	}
	maxId := tc.timeLine[len(tc.timeLine)-1].ID
	numId, _ := strconv.Atoi(string(maxId))

	tc.timeLine = append(tc.timeLine, &mastodon.Status{
		ID:        mastodon.ID(strconv.Itoa(numId + 1)),
		CreatedAt: time.Now(),
		Reblog:    &targetToot,
	})
}

func (tc *TestClient) RemoveMyMessageByID(id mastodon.ID) {
	for i, toot := range tc.timeLine {
		if toot.ID == id {
			tc.timeLine = append(tc.timeLine[:i], tc.timeLine[i+1:]...)
		}
	}
}

func (tc *TestClient) UpdateAvatar(filename string) {
}
