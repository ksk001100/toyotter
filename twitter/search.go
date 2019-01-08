package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// SearchTweet ツイート検索
func SearchTweet(api *anaconda.TwitterApi, text string, v url.Values) {
	searchResult, err := api.GetSearch(text, v)

	if err != nil {
		modules.ErrorMessage("ツイート検索に失敗したよ")
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}

// SearchUser ユーザー検索
func SearchUser(api *anaconda.TwitterApi, text string, v url.Values) {
	searchResult, err := api.GetUserSearch(text, v)

	if err != nil {
		modules.ErrorMessage("ツイート検索に失敗したよ")
	}

	for _, user := range searchResult {
		fmt.Println(modules.GetFormatUser(user))
	}
}
