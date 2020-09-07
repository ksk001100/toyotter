package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli"
)

// ListCommand list command function
func ListCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "list",
		Aliases: []string{"li"},
		Usage:   "toyotter list [...option]",
		Flags:   listFlags(),
		Action:  listAction,
	}
}

func listFlags() []cli.Flag {
	countFlag := cli.StringFlag{
		Name:  "count c",
		Value: "10",
		Usage: "toyotter list --count=[count]",
	}
	return []cli.Flag{
		&countFlag,
	}
}

func listAction(c *cli.Context) error {
	twitter.Lists(api, v)
	return nil
}
