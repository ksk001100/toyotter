package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/modules"
)

// Lists lists function
func Lists(api *anaconda.TwitterApi, v url.Values) {
	user, err := api.GetSelf(v)
	lists, err := api.GetListsOwnedBy(user.Id, v)

	if err != nil {
		modules.ErrorMessage("Get lists failed")
	}

	for _, list := range lists {
		fmt.Println(modules.GetFormatList(list))
	}
}
