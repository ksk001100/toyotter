package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli/v2"
)

// FollowCommand follow command function
func FollowCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "follow",
		Aliases: []string{"flw"},
		Usage:   "toyotter follow [screenName]",
		Flags:   followFlags(),
		Action:  followAction,
	}
}

func followFlags() []cli.Flag {
	deleteFlag := cli.BoolFlag{
		Name:  "delete, del, d",
		Usage: "toyotter follow [screenName] --delete",
	}
	return []cli.Flag{
		&deleteFlag,
	}
}

func followAction(c *cli.Context) error {
	screenName := c.Args().First()
	if c.Bool("delete") {
		twitter.UnFollow(api, screenName)
	} else {
		twitter.Follow(api, screenName)
	}
	return nil
}
