package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Follow フォロー
func Follow(api *anaconda.TwitterApi, screenName string) {
	user, err := api.FollowUser(screenName)

	if err != nil {
		modules.ErrorMessage("フォローに失敗したよ")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// UnFollow フォロー解除
func UnFollow(api *anaconda.TwitterApi, screenName string) {
	user, err := api.UnfollowUser(screenName)

	if err != nil {
		modules.ErrorMessage("フォロー解除に失敗したよ")
	}

	fmt.Println(modules.GetFormatUser(user))
}
