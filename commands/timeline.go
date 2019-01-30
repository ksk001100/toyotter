package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// TimelineCommand timeline command function
func TimelineCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Usage:   "toyotter2 timeline [...option]",
		Flags:   timelineFlags(),
		Action:  timelineAction,
	}
}

func timelineFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "count, c",
			Value: "5",
			Usage: "toyotter2 timeline --count=[count]",
		},
	}
}

func timelineAction(c *cli.Context) error {
	v.Set("count", c.String("count"))
	twitter.HomeTimeline(api, v)
	return nil
}
