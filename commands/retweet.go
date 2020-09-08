package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli/v2"
)

// RetweetCommand retweet command function
func RetweetCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "retweet",
		Aliases: []string{"rt"},
		Usage:   "toyotter retweet [tweetID]",
		Flags:   retweetFlags(),
		Action:  retweetAction,
	}
}

func retweetFlags() []cli.Flag {
	deleteFlag := cli.BoolFlag{
		Name:  "delete, del, d",
		Usage: "toyotter retweet [tweetID] --delete",
	}
	return []cli.Flag{
		&deleteFlag,
	}
}

func retweetAction(c *cli.Context) error {
	tweetID, _ := strconv.ParseInt(c.Args().First(), 10, 64)
	if c.Bool("delete") {
		twitter.UnRetweet(api, tweetID)
	} else {
		twitter.Retweet(api, tweetID)
	}
	return nil
}
