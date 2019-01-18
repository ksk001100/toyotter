package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Mute ミュート
func Mute(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.MuteUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("ミュートに失敗したよ")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// UnMute ミュート解除
func UnMute(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.UnmuteUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("ミュート解除に失敗したよ")
	}

	fmt.Println(modules.GetFormatUser(user))
}
