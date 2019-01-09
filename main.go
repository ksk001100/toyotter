package main

import (
	"fmt"
	"net/url"
	"os"
	"os/user"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter2/modules"
	"github.com/KeisukeToyota/toyotter2/twitter"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func loadEnv() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(user.HomeDir + "/.env.toyotter")
	if err != nil {
		modules.ErrorMessage(".env.toyotterが見つからないよ")
	}
}

func getTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func main() {
	loadEnv()
	api := getTwitterAPI()
	v := url.Values{}

	app := cli.NewApp()
	app.Name = "toyotter2"
	app.Usage = "toyotter2 [command] [...option]"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "tweet",
			Aliases: []string{"tw"},
			Usage:   "toyotter2 tweet [text]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image, img",
					Usage: "toyotter2 tweet [text] --image=[imagePath]",
				},
			},
			Action: func(c *cli.Context) error {
				text := c.Args().First()
				if len(c.String("image")) > 0 {
					fmt.Println(c.String("image"))
					media := twitter.UploadMedia(api, c.String("image"))
					v.Set("media_ids", media.MediaIDString)
				}
				twitter.Tweet(api, text, v)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Usage:   "toyotter2 tweet delete [tweetID]",
					Action: func(c *cli.Context) error {
						tweetID, _ := strconv.ParseInt(c.Args().First(), 10, 64)
						twitter.DeleteTweet(api, tweetID)
						return nil
					},
				},
			},
		},
		{
			Name:    "timeline",
			Aliases: []string{"tl"},
			Usage:   "toyotter2 timeline",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "count, c",
					Value: "10",
					Usage: "toyotter2 timeline --count=[count]",
				},
			},
			Action: func(c *cli.Context) error {
				v.Set("count", c.String("count"))
				twitter.HomeTimeline(api, v)
				return nil
			},
		},
		{
			Name:  "search",
			Usage: "toyotter2 search [option]=[text]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "toyotter2 search --user=[text]",
				},
				cli.StringFlag{
					Name:  "tweet, tw",
					Usage: "toyotter2 search --tweet=[text]",
				},
				cli.StringFlag{
					Name:  "count, c",
					Value: "10",
					Usage: "toyotter2 search --{user|tweet}=[text] --count=[count]",
				},
			},
			Action: func(c *cli.Context) error {
				v.Set("count", c.String("count"))
				if len(c.String("tweet")) > 0 {
					text := c.String("tweet")
					twitter.SearchTweet(api, text, v)
				} else if len(c.String("user")) > 0 {
					text := c.String("user")
					twitter.SearchUser(api, text, v)
				}
				return nil
			},
		},
		{
			Name:    "retweet",
			Aliases: []string{"rt"},
			Usage:   "toyotter2 retweet [tweetID]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "delete, d",
					Usage: "toyotter2 retweet [tweetID] --delete",
				},
			},
			Action: func(c *cli.Context) error {
				tweetID, _ := strconv.ParseInt(c.Args().First(), 10, 64)
				if c.Bool("delete") {
					twitter.UnRetweet(api, tweetID)
				} else {
					twitter.Retweet(api, tweetID)
				}
				return nil
			},
		},
		{
			Name:    "favorite",
			Aliases: []string{"fav"},
			Usage:   "toyotter2 favorite [tweetID]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "delete, d",
					Usage: "toyotter2 favorite [tweetID] --delete",
				},
			},
			Action: func(c *cli.Context) error {
				tweetID, _ := strconv.ParseInt(c.Args().First(), 10, 64)
				if c.Bool("delete") {
					twitter.UnFavorite(api, tweetID)
				} else {
					twitter.Favorite(api, tweetID)
				}
				return nil
			},
		},
		{
			Name:    "follow",
			Aliases: []string{"flw"},
			Usage:   "toyotter2 follow [screenName]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "delete, d",
					Usage: "toyotter2 follow [screenName] --delete",
				},
			},
			Action: func(c *cli.Context) error {
				screenName := c.Args().First()
				if c.Bool("delete") {
					twitter.UnFollow(api, screenName)
				} else {
					twitter.Follow(api, screenName)
				}
				return nil
			},
		},
		{
			Name:    "block",
			Aliases: []string{"blk"},
			Usage:   "toyotter2 block [screenName]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "delete, d",
					Usage: "toyotter2 block [screenName] --delete",
				},
			},
			Action: func(c *cli.Context) error {
				screenName := c.Args().First()
				if c.Bool("delete") {
					twitter.UnBlock(api, screenName, v)
				} else {
					twitter.Block(api, screenName, v)
				}
				return nil
			},
		},
		// {
		// 	Name:  "dm",
		// 	Usage: "toyotter2 dm [screenName] [text]",
		// 	Action: func(c *cli.Context) error {
		// 		screenName := c.Args().Get(0)
		// 		text := c.Args().Get(1)
		// 		twitter.DirectMessage(api, screenName, text)
		// 		return nil
		// 	},
		// },
	}

	app.Run(os.Args)
}
