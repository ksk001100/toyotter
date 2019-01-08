package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// HomeTimeline タイムライン出力
func HomeTimeline(api *anaconda.TwitterApi, v url.Values) {
	tweets, err := api.GetHomeTimeline(v)

	if err != nil {
		modules.ErrorMessage("ツイート取得に失敗したよ")
	}

	for _, tweet := range tweets {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}
