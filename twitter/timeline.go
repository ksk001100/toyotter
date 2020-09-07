package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/modules"
)

// HomeTimeline hometimeline function
func HomeTimeline(api *anaconda.TwitterApi, v url.Values) {
	tweets, err := api.GetHomeTimeline(v)

	if err != nil {
		modules.ErrorMessage("Get home timeline failed")
	}

	for _, tweet := range tweets {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}

// ListTimeline List timeline function
func ListTimeline(api *anaconda.TwitterApi, listID int64, v url.Values) {
	tweets, err := api.GetListTweets(listID, false, v)

	if err != nil {
		modules.ErrorMessage("Get list timeline failed")
	}

	for _, tweet := range tweets {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}
