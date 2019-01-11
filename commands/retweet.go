package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// RetweetCommand リツイートコマンド
func RetweetCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "retweet"
	command.Aliases = []string{"rt"}
	command.Usage = "toyotter2 retweet [tweetID]"
	command.Flags = retweetFlags()
	command.Action = retweetAction
	return command
}

func retweetFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter2 retweet [tweetID] --delete",
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
