package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// MentionCommand mention command function
func MentionCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "mention",
		Aliases: []string{"men"},
		Usage:   "toyotter2 mension [option]",
		Flags:   mentionFlags(),
		Action:  mentionAction,
	}
}

func mentionFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "count, c",
			Value: "5",
			Usage: "toyotter2 mension --count 20",
		},
	}
}

func mentionAction(c *cli.Context) error {
	v.Set("count", c.String("count"))
	twitter.Mention(api, v)
	return nil
}
