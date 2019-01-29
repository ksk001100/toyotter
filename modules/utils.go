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

// GetFormatTweet ツイートの整形
func GetFormatTweet(tweet anaconda.Tweet) string {
	t, _ := tweet.CreatedAtTime()
	return fmt.Sprintf(getFormatTweetTemplate(),
		aurora.Brown(getJapanDateTimeString(t)).String(), aurora.Green(tweet.User.Name).String(),
		aurora.Green(tweet.User.ScreenName).String(), aurora.Red(tweet.User.IdStr).String(),
		aurora.Green(tweet.FavoriteCount).String(), aurora.Green(tweet.RetweetCount).String(),
		aurora.Bold(tweet.FullText).String(), aurora.Red(tweet.IdStr).String(),
		SeparatorString(),
	)
}

// GetFormatUser ユーザー情報整形
func GetFormatUser(user anaconda.User) string {
	return fmt.Sprintf(getFormatUserTemplate(),
		aurora.Green(user.Name).String(), aurora.Green(user.ScreenName).String(),
		aurora.Red(user.IdStr).String(), aurora.Green(user.FriendsCount).String(),
		aurora.Green(user.FollowersCount).String(), aurora.Bold(user.Description).String(),
		SeparatorString(),
	)
}

// ErrorMessage エラーメッセージ
func ErrorMessage(text string) {
	log.Fatal(aurora.Red(text))
}

// SeparatorString セパレーター文字列
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
[%s @%s | ユーザーID : %s]
[いいね数 : %s | リツイート数 : %s]

%s

[ツイートID : %s]

%s`
}

func getFormatUserTemplate() string {
	return `
[%s @%s | ユーザーID : %s]
[フォロー数 : %s | フォロワー数 : %s]

%s

%s`
}
