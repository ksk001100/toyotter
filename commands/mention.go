package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// MentionCommand メンションコマンド
func MentionCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "mention"
	command.Aliases = []string{"men"}
	command.Usage = "toyotter2 mension [option]"
	command.Flags = mentionFlags()
	command.Action = mentionAction
	return command
}

func mentionFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "count, c",
			Value: "10",
			Usage: "toyotter2 mension --count 20",
		},
	}
}

func mentionAction(c *cli.Context) error {
	v.Set("count", c.String("count"))
	twitter.Mention(api, v)
	return nil
}
