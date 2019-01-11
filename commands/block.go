package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// BlockCommand ブロックコマンド
func BlockCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "block"
	command.Aliases = []string{"blk"}
	command.Usage = "toyotter2 block [screenName]"
	command.Flags = blockFlags()
	command.Action = blockAction
	return command
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
