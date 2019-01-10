package twitter

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Tweet ツイート
func Tweet(api *anaconda.TwitterApi, text string, v url.Values) {
	tweet, err := api.PostTweet(text, v)

	if err != nil {
		modules.ErrorMessage("ツイートに失敗したよ")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// DeleteTweet ツイート削除
func DeleteTweet(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.DeleteTweet(tweetID, false)

	if err != nil {
		modules.ErrorMessage("ツイート削除に失敗したよ")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// Reply リプライ
func Reply(api *anaconda.TwitterApi, text string, tweetID int64) {
	r := regexp.MustCompile(`@[0-9a-zA-Z_]{1,15}`)
	tweet, err := api.GetTweet(tweetID, url.Values{})
	if err != nil {
		modules.ErrorMessage("ツイートが見つからないよ")
	}

	r.FindAllStringSubmatch(tweet.FullText, -1)
}
