package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli"
)

// FavoriteCommand favorite command function
func FavoriteCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "favorite",
		Aliases: []string{"fav"},
		Usage:   "toyotter favorite [tweetID]",
		Flags:   favoriteFlags(),
		Action:  favoriteAction,
	}
}

func favoriteFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter favorite [tweetID] --delete",
		},
	}
}

func favoriteAction(c *cli.Context) error {
	tweetID, _ := strconv.ParseInt(c.Args().First(), 10, 64)
	if c.Bool("delete") {
		twitter.UnFavorite(api, tweetID)
	} else {
		twitter.Favorite(api, tweetID)
	}
	return nil
}
