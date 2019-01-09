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
