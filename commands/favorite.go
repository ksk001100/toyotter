package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// FavoriteCommand いいねコマンド
func FavoriteCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "favorite"
	command.Aliases = []string{"fav"}
	command.Usage = "toyotter2 favorite [tweetID]"
	command.Flags = favoriteFlags()
	command.Action = favoriteAction
	return command
}

func favoriteFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter2 favorite [tweetID] --delete",
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