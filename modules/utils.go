package modules

import (
	"fmt"
	"log"

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
	return fmt.Sprintf("\n[%s] (@%s) %s\n%s [Tweet ID : %s]",
		aurora.Brown(tweet.CreatedAt).String(),
		aurora.Magenta(tweet.User.ScreenName).String(),
		aurora.Green(tweet.User.Name).String(),
		aurora.Bold(tweet.FullText).String(),
		aurora.Red(tweet.IdStr).String())
}

// GetFormatUser ユーザー情報整形
func GetFormatUser(user anaconda.User) string {
	return fmt.Sprintf("\n(@%s) %s [User ID : %s]",
		aurora.Magenta(user.ScreenName).String(),
		aurora.Green(user.Name).String(),
		aurora.Red(user.IdStr).String())
}

// ErrorMessage エラーメッセージ
func ErrorMessage(text string) {
	log.Fatal(aurora.Red(text))
}
