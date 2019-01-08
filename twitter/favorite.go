package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Favorite いいね
func Favorite(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.Favorite(tweetID)

	if err != nil {
		modules.ErrorMessage("いいねに失敗したよ")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// UnFavorite いいね解除
func UnFavorite(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.Unfavorite(tweetID)

	if err != nil {
		modules.ErrorMessage("いいね解除に失敗したよ")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}
