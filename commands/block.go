package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// BlockCommand block command function
func BlockCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "block",
		Aliases: []string{"blk"},
		Usage:   "toyotter2 block [screenName]",
		Flags:   blockFlags(),
		Action:  blockAction,
	}
}

func blockFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "delete, del, d",
			Usage: "toyotter2 block [screenName] --delete",
		},
		cli.BoolFlag{
			Name:  "list, l",
			Usage: "toyotter2 block --list",
		},
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
