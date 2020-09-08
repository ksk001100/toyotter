package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli/v2"
)

// BlockCommand block command function
func BlockCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "block",
		Aliases: []string{"blk"},
		Usage:   "toyotter block [screenName]",
		Flags:   blockFlags(),
		Action:  blockAction,
	}
}

func blockFlags() []cli.Flag {
	deleteFlag := cli.BoolFlag{
		Name:  "delete, del, d",
		Usage: "toyotter block [screenName] --delete",
	}
	listFlag := cli.BoolFlag{
		Name:  "list, l",
		Usage: "toyotter block --list",
	}
	return []cli.Flag{
		&deleteFlag,
		&listFlag,
	}
}

func blockAction(c *cli.Context) error {
	if c.Bool("list") {
		twitter.BlockUser(api, url.Values{})
	} else {
		screenName := c.Args().First()
		if c.Bool("delete") {
			twitter.UnBlock(api, screenName, v)
		} else {
			twitter.Block(api, screenName, v)
		}
	}
	return nil
}
