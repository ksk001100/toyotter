package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Favorite favorite function
func Favorite(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.Favorite(tweetID)

	if err != nil {
		modules.ErrorMessage("Favorite failed")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// UnFavorite unfavorite function
func UnFavorite(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.Unfavorite(tweetID)

	if err != nil {
		modules.ErrorMessage("UnFavorite failed")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}
