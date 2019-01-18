package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// MuteCommand ミュートコマンド
func MuteCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "mute",
		Aliases: []string{"mu"},
		Usage:   "toyotter2 mute [screenName]",
		Flags:   muteFlags(),
		Action:  muteAction,
	}
}

func muteFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter2 mute [screenName] --delete",
		},
	}
}

func muteAction(c *cli.Context) error {
	screenName := c.Args().First()
	if c.Bool("delete") {
		twitter.UnBlock(api, screenName, v)
	} else {
		twitter.Block(api, screenName, v)
	}
	return nil
}
