package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Mension メンション
func Mension(api *anaconda.TwitterApi, v url.Values) {
	tweets, err := api.GetMentionsTimeline(v)

	if err != nil {
		modules.ErrorMessage("メンション取得に失敗したよ")
	}

	for _, tweet := range tweets {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}
