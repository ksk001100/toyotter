package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli/v2"
)

// MentionCommand mention command function
func MentionCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "mention",
		Aliases: []string{"men"},
		Usage:   "toyotter mension [option]",
		Flags:   mentionFlags(),
		Action:  mentionAction,
	}
}

func mentionFlags() []cli.Flag {
	countFlag := cli.StringFlag{
		Name:  "count, c",
		Value: "5",
		Usage: "toyotter mension --count 20",
	}
	return []cli.Flag{
		&countFlag,
	}
}

func mentionAction(c *cli.Context) error {
	v.Set("count", c.String("count"))
	twitter.Mention(api, v)
	return nil
}
