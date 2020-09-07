package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/modules"
)

// Mention mention function
func Mention(api *anaconda.TwitterApi, v url.Values) {
	tweets, err := api.GetMentionsTimeline(v)

	if err != nil {
		modules.ErrorMessage("Get mentions failed")
	}

	for _, tweet := range tweets {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}
