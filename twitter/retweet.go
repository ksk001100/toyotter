package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/modules"
)

// Retweet retweet function
func Retweet(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.Retweet(tweetID, false)

	if err != nil {
		modules.ErrorMessage("Retweet failed")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// UnRetweet unretweet function
func UnRetweet(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.UnRetweet(tweetID, false)

	if err != nil {
		modules.ErrorMessage("UnRetweet failed")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}
