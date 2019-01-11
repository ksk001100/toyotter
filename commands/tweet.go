package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/urfave/cli"
)

// TweetCommand ツイートコマンド
func TweetCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	command := cli.Command{}
	command.Name = "tweet"
	command.Aliases = []string{"tw"}
	command.Usage = "toyotter2 tweet [text]"
	command.Flags = tweetFlags()
	command.Action = tweetAction
	return command
}

func tweetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "image, img",
			Usage: "toyotter2 tweet [text] --image=[imagePath]",
		},
		cli.StringFlag{
			Name:  "reply, rep",
			Usage: "toyotter2 tweet [text] --reply=[tweetID]",
		},
		cli.StringFlag{
			Name:  "delete, del, d",
			Usage: "toyotter2 tweet --delete=[tweetID]",
		},
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
		} else {
			twitter.Tweet(api, text, v)
		}
	}
	return nil
}
