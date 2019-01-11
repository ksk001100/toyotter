package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// Block ブロック
func Block(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.BlockUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("ブロックに失敗したよ")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// UnBlock ブロック解除
func UnBlock(api *anaconda.TwitterApi, screenName string, v url.Values) {
	user, err := api.UnblockUser(screenName, v)

	if err != nil {
		modules.ErrorMessage("ブロック解除に失敗したよ")
	}

	fmt.Println(modules.GetFormatUser(user))
}

// BlockUser ブロックユーザー一覧
func BlockUser(api *anaconda.TwitterApi, v url.Values) {
	userCursor, err := api.GetBlocksList(v)

	if err != nil {
		modules.ErrorMessage("ブロックユーザー一覧の取得に失敗したよ")
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
