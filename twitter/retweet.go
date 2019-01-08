package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Retweet リツイート
func Retweet(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.Retweet(tweetID, false)

	if err != nil {
		modules.ErrorMessage("リツイートに失敗したよ")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// UnRetweet リツイート解除
func UnRetweet(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.UnRetweet(tweetID, false)

	if err != nil {
		modules.ErrorMessage("リツイート解除に失敗したよ")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}
