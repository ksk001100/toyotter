package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/twitter"
	"github.com/urfave/cli"
)

// SearchCommand search command function
func SearchCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "toyotter search [option]=[text]",
		Flags:   searchFlags(),
		Action:  searchAction,
	}
}

func searchFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Usage: "toyotter search --user=[text]",
		},
		cli.StringFlag{
			Name:  "tweet, tw",
			Usage: "toyotter search --tweet=[text]",
		},
		cli.StringFlag{
			Name:  "count, c",
			Value: "5",
			Usage: "toyotter search --{user|tweet}=[text] --count=[count]",
		},
	}
}

func searchAction(c *cli.Context) error {
	v.Set("count", c.String("count"))
	if len(c.String("tweet")) > 0 {
		text := c.String("tweet")
		twitter.SearchTweet(api, text, v)
	} else if len(c.String("user")) > 0 {
		text := c.String("user")
		twitter.SearchUser(api, text, v)
	}
	return nil
}
