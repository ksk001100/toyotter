package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// TimelineCommand タイムラインコマンド
func TimelineCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "timeline"
	command.Aliases = []string{"tl"}
	command.Usage = "toyotter2 timeline"
	command.Flags = timelineFlags()
	command.Action = timelineAction
	return command
}

func timelineFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "count, c",
			Value: "10",
			Usage: "toyotter2 timeline --count=[count]",
		},
	}
}

func timelineAction(c *cli.Context) error {
	v.Set("count", c.String("count"))
	twitter.HomeTimeline(api, v)
	return nil
}
