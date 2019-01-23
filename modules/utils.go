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

// ContainsString 配列に特定の文字列が含まれるかチェック
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetFormatTweet ツイートの整形
func GetFormatTweet(tweet anaconda.Tweet) string {
	width, _, _ := terminal.GetSize(0)
	t, _ := tweet.CreatedAtTime()
	japanTime, _ := time.LoadLocation("Asia/Tokyo")
	createdAt := t.In(japanTime)
	datetime := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		createdAt.Year(), createdAt.Month(), createdAt.Day(),
		createdAt.Hour(), createdAt.Minute(), createdAt.Second(),
	)
	return fmt.Sprintf("\n[%s]\n[アカウント名 : @%s | 名前 : %s | ユーザーID : %s]\n[いいね数 : %s | リツイート数 : %s]\n\n%s\n\n[ツイートID : %s]\n\n%s\n",
		aurora.Brown(datetime).String(), aurora.Green(tweet.User.ScreenName).String(),
		aurora.Green(tweet.User.Name).String(), aurora.Red(tweet.User.IdStr),
		aurora.Green(tweet.FavoriteCount).String(), aurora.Green(tweet.RetweetCount).String(),
		aurora.Bold(tweet.FullText).String(), aurora.Red(tweet.IdStr).String(),
		strings.Repeat("-", width),
	)
}

// GetFormatUser ユーザー情報整形
func GetFormatUser(user anaconda.User) string {
	return fmt.Sprintf("\n(@%s) %s [ユーザーID : %s]",
		aurora.Magenta(user.ScreenName).String(),
		aurora.Green(user.Name).String(),
		aurora.Red(user.IdStr).String(),
	)
}

// ErrorMessage エラーメッセージ
func ErrorMessage(text string) {
	log.Fatal(aurora.Red(text))
}
