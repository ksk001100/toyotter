package modules

import (
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/ChimeraCoder/anaconda"
	"github.com/logrusorgru/aurora"
)

// GetFormatTweet format tweet function
func GetFormatTweet(tweet anaconda.Tweet) string {
	t, _ := tweet.CreatedAtTime()
	return fmt.Sprintf(getFormatTweetTemplate(),
		aurora.Brown(getJapanDateTimeString(t)).String(), aurora.Green(tweet.User.Name).String(),
		aurora.Green(tweet.User.ScreenName).String(), aurora.Red(tweet.User.IdStr).String(),
		aurora.Green(tweet.FavoriteCount).String(), aurora.Green(tweet.RetweetCount).String(),
		aurora.Bold(tweet.FullText).String(), aurora.Red(tweet.IdStr).String(),
		aurora.Green(getTweetURL(tweet)),
		SeparatorString(),
	)
}

// GetFormatUser format user information function
func GetFormatUser(user anaconda.User) string {
	return fmt.Sprintf(getFormatUserTemplate(),
		aurora.Green(user.Name).String(), aurora.Green(user.ScreenName).String(),
		aurora.Red(user.IdStr).String(), aurora.Green(user.FriendsCount).String(),
		aurora.Green(user.FollowersCount).String(), aurora.Bold(user.Description).String(),
		SeparatorString(),
	)
}

// GetFormatList format list information function
func GetFormatList(list anaconda.List) string {
	return fmt.Sprintf(getFormatListTemplate(),
		aurora.Red(list.Id),
		aurora.Blue(list.Name),
		aurora.Red(list.Description),
		aurora.Yellow(list.MemberCount),
		aurora.Green(getListURL(list)),
		SeparatorString(),
	)
}

// ErrorMessage error message function
func ErrorMessage(text string) {
	log.Fatal(aurora.Red(text))
}

// SeparatorString separator function
func SeparatorString() string {
	width, _, _ := terminal.GetSize(0)
	return strings.Repeat("-", width)
}

func getJapanDateTimeString(t time.Time) string {
	japanTime, _ := time.LoadLocation("Asia/Tokyo")
	createdAt := t.In(japanTime)
	return fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		createdAt.Year(), createdAt.Month(), createdAt.Day(),
		createdAt.Hour(), createdAt.Minute(), createdAt.Second(),
	)
}

func getFormatTweetTemplate() string {
	return `
[%s]
[%s @%s | UserID : %s]
[Favorite : %s | Retweet : %s]

%s

[TweetID : %s]
[URL : %s]

%s`
}

func getFormatUserTemplate() string {
	return `
[%s @%s | UserID : %s]
[Followed : %s | Follower : %s]

%s

%s`
}

func getFormatListTemplate() string {
	return `
[ID : %d]
[Name : %s]
[Description : %s]
[Member count : %d]
[URL : %s]

%s`
}

func getTweetURL(tweet anaconda.Tweet) string {
	return fmt.Sprintf("https://twitter.com/%s/status/%s", tweet.User.ScreenName, tweet.IdStr)
}

func getListURL(list anaconda.List) string {
	return fmt.Sprintf("https://twitter.com%s", list.URL)
}


