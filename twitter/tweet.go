package twitter

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/modules"
)

// Tweet tweet function
func Tweet(api *anaconda.TwitterApi, text string, v url.Values) {
	tweet, err := api.PostTweet(text, v)

	if err != nil {
		modules.ErrorMessage("Tweet failed")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// DeleteTweet delete tweet function
func DeleteTweet(api *anaconda.TwitterApi, tweetID int64) {
	tweet, err := api.DeleteTweet(tweetID, false)

	if err != nil {
		modules.ErrorMessage("Delte tweet failed")
	}

	fmt.Println(modules.GetFormatTweet(tweet))
}

// Reply reply function
func Reply(api *anaconda.TwitterApi, text string, tweetID int64, v url.Values) {
	v.Set("in_reply_to_status_id", strconv.FormatInt(tweetID, 10))

	tweet, err := api.GetTweet(tweetID, url.Values{})
	if err != nil {
		modules.ErrorMessage("Get tweet failed")
	}

	replayUserName := "@" + tweet.User.ScreenName + " "
	replayTweet, replyErr := api.PostTweet(replayUserName+text, v)
	if replyErr != nil {
		modules.ErrorMessage("Reply failed")
	}

	fmt.Println(modules.GetFormatTweet(replayTweet))
}

// QuoteTweet quote tweet function
func QuoteTweet(api *anaconda.TwitterApi, text string, quoteTweetID int64, v url.Values) {
	quoteTweet, err := api.GetTweet(quoteTweetID, url.Values{})

	if err != nil {
		modules.ErrorMessage("The tweet you quoted is not valid")
	}

	quoteTweetURL := modules.GetTweetURL(quoteTweet)
	quotedTweetText := fmt.Sprintf("%s\n%s", text, quoteTweetURL)

	Tweet(api, quotedTweetText, v)
}
