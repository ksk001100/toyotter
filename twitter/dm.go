	package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
)

// DirectMessage DM送信
func DirectMessage(api *anaconda.TwitterApi, screenName string, text string) {
	msg, err := api.PostDMToScreenName(text, screenName)

	if err != nil {
		modules.ErrorMessage("DM送信に失敗したよ")
	}

	fmt.Println(msg.Text)
}
