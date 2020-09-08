package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli/v2"
)

// TweetCommand tweet command function
func TweetCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "tweet",
		Aliases: []string{"tw"},
		Usage:   "toyotter tweet [text] [...option]",
		Flags:   tweetFlags(),
		Action:  tweetAction,
	}
}

func tweetFlags() []cli.Flag {
	imageFlag := cli.StringFlag{
			Name:  "image, img",
			Usage: "toyotter tweet [text] --image=[imagePath]",
		}
	replyFlag := cli.StringFlag{
			Name:  "reply, rep",
			Usage: "toyotter tweet [text] --reply=[tweetID]",
		}
	deleteFlag := cli.StringFlag{
			Name:  "delete, del, d",
			Usage: "toyotter tweet --delete=[tweetID]",
		}
	quoteFlag := cli.StringFlag{
			Name:  "quote, q",
			Usage: "toyotter tweet [text] --quote=[tweetID]",
		}
	return []cli.Flag{
		&imageFlag,
		&replyFlag,
		&deleteFlag,
		&quoteFlag,
	}
}

func tweetAction(c *cli.Context) error {
	if len(c.String("delete")) > 0 {
		tweetID, _ := strconv.ParseInt(c.String("delete"), 10, 64)
		twitter.DeleteTweet(api, tweetID)
	} else {
		text := c.Args().First()
		if len(c.String("image")) > 0 {
			media := twitter.UploadMedia(api, c.String("image"))
			v.Set("media_ids", media.MediaIDString)
		}

		if len(c.String("reply")) > 0 {
			tweetID, _ := strconv.ParseInt(c.String("reply"), 10, 64)
			twitter.Reply(api, text, tweetID, v)
		} else if len(c.String("quote")) > 0 {
			quoteTweetID, _ := strconv.ParseInt(c.String("quote"), 10, 64)
			twitter.QuoteTweet(api, text, quoteTweetID, v)
		} else {
			twitter.Tweet(api, text, v)
		}
	}
	return nil
}
