package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/modules"
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
