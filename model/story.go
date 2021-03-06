package model

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/SyntropyDev/milli"
	"github.com/SyntropyDev/sqlutil"
	"github.com/SyntropyDev/val"
	"github.com/coopernurse/gorp"
	"github.com/huandu/facebook"
	"github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
	"github.com/lann/squirrel"
)

const (
	ObjectNameStory = "Story"
	TableNameStory  = "stories"
)

type Story struct {
	ID      int64  `json:"id"`
	Created int64  `json:"created" val:"nonzero"`
	Updated int64  `json:"updated" val:"nonzero"`
	Deleted bool   `json:"deleted" merge:"true"`
	Object  string `db:"-" json:"object"`

	MemberID           int64     `json:"memberId" val:"nonzero"`
	MemberName         string    `json:"memberName"`
	FeedID             int64     `json:"feedId" val:"nonzero"`
	FeedIdentifier     string    `json:"feedIdentifier"`
	FeedType           string    `json:"feedType"`
	Timestamp          int64     `json:"timestamp"`
	Body               string    `json:"body"`
	SourceURL          string    `json:"sourceUrl"`
	SourceID           string    `json:"sourceId"`
	Score              float64   `json:"score"`
	Latitude           float64   `json:"-"`
	Longitude          float64   `json:"-"`
	LinksRaw           string    `json:"-"`
	ImagesRaw          string    `json:"-"`
	HashtagsRaw        string    `json:"-"`
	LastDecayTimestamp int64     `json:"-"`
	CategoryIds        []int64   `db:"-" json:"categoryIds"`
	Links              []string  `db:"-" json:"links"`
	Images             []string  `db:"-" json:"images"`
	Hashtags           []string  `db:"-" json:"hashTags"`
	Location           []float64 `db:"-" json:"location"`
	MemberIcon         string    `db:"-" json:"memberIcon"`
}

func NewFacebookStory(member *Member, feed *Feed, post *FacebookPost) *Story {
	t, err := time.Parse(time.RFC3339Nano, post.CreatedAt)
	if err != nil {
		t = time.Now()
	}

	// if message is blank, forget the post
	if post.Message == "" {
		return nil
	}

	urlParts := strings.Split(post.Link, "?")
	if len(urlParts) > 0 {
		post.Link = urlParts[0]
	}

	post.Picture = ""
	if post.Type == "photo" {
		session := facebookSession()
		route := fmt.Sprintf("/%s?fields=images", post.ObjectId)
		result, err := session.Api(route, facebook.GET, nil)
		if err == nil {
			image := &FacebookPhoto{}
			if err := result.Decode(image); err == nil {
				if len(image.Images) > 0 && !strings.Contains(image.Images[0].Source, "?") {
					post.Picture = image.Images[0].Source
				}
			}
		}
	}

	return &Story{
		MemberID:       member.ID,
		MemberName:     member.Name,
		FeedID:         feed.ID,
		FeedIdentifier: feed.Identifier,
		Timestamp:      milli.Timestamp(t),
		Body:           strings.TrimSpace(post.Message),
		FeedType:       string(FeedTypeFacebook),
		SourceURL:      post.Link,
		SourceID:       post.Id,
		Latitude:       0.0,
		Longitude:      0.0,
		Score:          2,
		LinksRaw:       "",
		HashtagsRaw:    "",
		ImagesRaw:      post.Picture,
	}
}

func NewStoryTwitter(member *Member, feed *Feed, tweet anaconda.Tweet) *Story {
	hashtags := []string{}
	for _, hashtag := range tweet.Entities.Hashtags {
		hashtags = append(hashtags, hashtag.Text)
	}

	urls := []string{}
	for _, url := range tweet.Entities.Urls {
		urls = append(urls, url.Url)
	}

	t, err := tweet.CreatedAtTime()
	if err != nil {
		t = time.Now()
	}

	images := []string{}
	for _, media := range tweet.Entities.Media {
		if media.Type == "image" {
			images = append(images, media.Media_url)
		}
	}
	score := tweet.FavoriteCount + (2 * tweet.RetweetCount)
	sourceURL := fmt.Sprintf("http://twitter.com/%s/status/%s", tweet.User.ScreenName, tweet.IdStr)
	return &Story{
		MemberID:       member.ID,
		MemberName:     member.Name,
		FeedID:         feed.ID,
		FeedIdentifier: feed.Identifier,
		Timestamp:      milli.Timestamp(t),
		Body:           tweet.Text,
		FeedType:       string(FeedTypeTwitter),
		SourceURL:      sourceURL,
		SourceID:       tweet.IdStr,
		Latitude:       0.0,
		Longitude:      0.0,
		Score:          float64(score),
		LinksRaw:       strings.Join(urls, ","),
		HashtagsRaw:    strings.Join(hashtags, ","),
		ImagesRaw:      strings.Join(images, ","),
	}
}

func NewStoryRSS(member *Member, feed *Feed, item *feeder.Item) *Story {
	// parse pub date
	itemTime, err := item.ParsedPubDate()
	if err != nil {
		itemTime = time.Now()
	}
	// form links
	links := []string{}
	for _, link := range item.Links {
		links = append(links, link.Href)
	}

	isAtom := func() bool {
		return item.Id != ""
	}

	sourceID := ""
	body := ""
	if isAtom() {
		sourceID = item.Id
		body = item.Content.Text
	} else { // is RSS
		sourceID = *item.Guid
		// use description or title for body
		body = item.Description
		if body == "" {
			body = item.Title
		}
	}

	// parse html for images
	images := []string{}
	doc := xmlx.New()
	doc.LoadString(strings.ToLower(body), nil)
	imgNodes := doc.SelectNodesRecursive("", "img")
	for _, img := range imgNodes {
		images = append(images, img.As("", "src"))
	}

	return &Story{
		MemberID:       member.ID,
		MemberName:     member.Name,
		FeedID:         feed.ID,
		FeedIdentifier: feed.Identifier,
		Timestamp:      milli.Timestamp(itemTime),
		Body:           body,
		FeedType:       string(FeedTypeRSS),
		SourceURL:      "",
		SourceID:       sourceID,
		Latitude:       0.0,
		Longitude:      0.0,
		LinksRaw:       strings.Join(links, ","),
		ImagesRaw:      strings.Join(images, ","),
	}
}

func DecayScores(s gorp.SqlExecutor) error {
	current := milli.Timestamp(time.Now())
	yesterday := milli.Timestamp(time.Now().Add(time.Hour * -24))
	tenDaysAgo := milli.Timestamp(time.Now().Add((time.Hour * 24) * 10))
	query := squirrel.Select("*").From(TableNameStory).
		Where("!(LastDecayTimestamp between ? and ?)", yesterday, current).
		Where("Timestamp > ?", tenDaysAgo)

	stories := []*Story{}
	sqlutil.Select(s, query, &stories)
	for _, story := range stories {
		story.Score /= 2.0
		story.LastDecayTimestamp = milli.Timestamp(time.Now())
		if _, err := s.Update(story); err != nil {
			return err
		}
	}
	return nil
}

func (story *Story) LinksSlice() []string {
	return sliceFromString(story.LinksRaw)
}

func (story *Story) ImagesSlice() []string {
	return sliceFromString(story.ImagesRaw)
}

func (story *Story) HashtagsSlice() []string {
	return sliceFromString(story.HashtagsRaw)
}

func (story *Story) LocationCoords() []float64 {
	return []float64{story.Latitude, story.Longitude}
}

func (story *Story) CalculateScore(s gorp.SqlExecutor) error {
	f := &Feed{}
	if err := sqlutil.SelectOneRelation(s, TableNameFeed, story.FeedID, f); err != nil {
		return err
	}

	score := 0.0
	if len(story.ImagesSlice()) > 0 {
		score += 10.0
	}
	if len(story.LinksSlice()) > 0 {
		score += 2.0
	}
	if story.Latitude != 0.0 {
		score += 10.0
	}

	switch FeedType(f.Type) {
	case FeedTypeFacebook:
		score += 3.0
	case FeedTypeTwitter:
		score += 3.0
	}

	// randomize score
	score += (10.0 * rand.Float64())

	// increase score for timely posts
	dur := milli.Time(story.Timestamp).Sub(milli.Time(f.LastRetrieved))
	durScore := float64(dur / time.Hour)
	if durScore > 12.0 {
		durScore = 12.0
	} else if durScore < 0.0 {
		durScore = 0.0
	}
	score += durScore

	story.Score += score

	return nil
}

func (story *Story) Validate() error {
	if valid, errMap := val.Struct(story); !valid {
		return ErrorFromMap(errMap)
	}
	return nil
}

func (story *Story) PreInsert(s gorp.SqlExecutor) error {
	story.Created = milli.Timestamp(time.Now())
	story.Updated = milli.Timestamp(time.Now())
	story.LastDecayTimestamp = milli.Timestamp(time.Now())
	story.CalculateScore(s)
	return story.Validate()
}

func (story *Story) PreUpdate(s gorp.SqlExecutor) error {
	story.Updated = milli.Timestamp(time.Now())
	return story.Validate()
}

func (story *Story) PostInsert(s gorp.SqlExecutor) error {
	m := &Member{}
	if err := sqlutil.SelectOneRelation(s, TableNameMember, story.MemberID, m); err != nil {
		return err
	}

	images := append(story.ImagesSlice(), m.ImagesSlice()...)
	m.SetImages(images)

	hashtags := append(story.HashtagsSlice(), m.HashtagsSlice()...)
	m.SetHashtags(hashtags)

	s.Update(m)

	feed := &Feed{}
	if err := sqlutil.SelectOneRelation(s, TableNameFeed, story.FeedID, feed); err != nil {
		return err
	}
	if story.Timestamp > feed.LastRetrieved {
		feed.LastRetrieved = story.Timestamp
		s.Update(feed)
	}
	return nil
}

func (story *Story) PostGet(s gorp.SqlExecutor) error {
	story.Object = ObjectNameStory
	story.Links = story.LinksSlice()
	story.Images = story.ImagesSlice()
	story.Hashtags = story.HashtagsSlice()
	story.Location = story.LocationCoords()

	m := &Member{}
	if err := sqlutil.SelectOneRelation(s, TableNameMember, story.MemberID, m); err != nil {
		return err
	}

	story.MemberIcon = m.Icon
	story.CategoryIds = m.CategoryIds
	return nil
}

// CrudResource interface

func (story *Story) TableName() string {
	return TableNameStory
}

func (story *Story) TableId() int64 {
	return story.ID
}

func (story *Story) Delete() {
	story.Deleted = true
}

func sliceFromString(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, ",")
}
