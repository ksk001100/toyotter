package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Mute mute function
func Mute(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.MuteUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("Mute failed")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// UnMute unmute function
func UnMute(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.UnmuteUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("UnMute failed")
	}

	fmt.Println(modules.GetFormatUser(user))
}
