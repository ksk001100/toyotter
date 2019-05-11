package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/twitter"
	"github.com/urfave/cli"
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
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter retweet [tweetID] --delete",
		},
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
