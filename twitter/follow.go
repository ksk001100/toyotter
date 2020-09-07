package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/modules"
)

// Follow follow function
func Follow(api *anaconda.TwitterApi, screenName string) {
	user, err := api.FollowUser(screenName)

	if err != nil {
		modules.ErrorMessage("Follow failed")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// UnFollow unfollow function
func UnFollow(api *anaconda.TwitterApi, screenName string) {
	user, err := api.UnfollowUser(screenName)

	if err != nil {
		modules.ErrorMessage("UnFollow failed")
	}

	fmt.Println(modules.GetFormatUser(user))
}
