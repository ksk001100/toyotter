package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// SearchTweet search tweet function
func SearchTweet(api *anaconda.TwitterApi, text string, v url.Values) {
	searchResult, err := api.GetSearch(text, v)

	if err != nil {
		modules.ErrorMessage("Search tweet failed")
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(modules.GetFormatTweet(tweet))
	}
}

// SearchUser search user function
func SearchUser(api *anaconda.TwitterApi, text string, v url.Values) {
	searchResult, err := api.GetUserSearch(text, v)

	if err != nil {
		modules.ErrorMessage("Search user failedf")
	}

	for _, user := range searchResult {
		fmt.Println(modules.GetFormatUser(user))
	}
}
