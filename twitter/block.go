package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/modules"
)

// Block block function
func Block(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.BlockUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("Block failed")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// UnBlock unblock function
func UnBlock(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.UnblockUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("UnBlock failed")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// BlockUser block user list function
func BlockUser(api *anaconda.TwitterApi, v url.Values) {
	userCursor, err := api.GetBlocksList(v)

	if err != nil {
		modules.ErrorMessage("Get block user list failed")
	}

	for _, user := range userCursor.Users {
		fmt.Println(modules.GetFormatUser(user))
	}

	for userCursor.Next_cursor != 0 {
		v.Set("cursor", userCursor.Next_cursor_str)
		userCursor, _ = api.GetBlocksList(v)
		for _, user := range userCursor.Users {
			fmt.Println(modules.GetFormatUser(user))
		}
	}
}
