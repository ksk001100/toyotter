package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// FollowCommand フォローコマンド
func FollowCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "follow"
	command.Aliases = []string{"flw"}
	command.Usage = "toyotter2 follow [screenName]"
	command.Flags = followFlags()
	command.Action = followAction
	return command
}

func followFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter2 follow [screenName] --delete",
		},
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
