package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"os/user"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"github.com/ksk001100/toyotter/commands"
	"github.com/ksk001100/toyotter/modules"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

func loadEnv() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(u.HomeDir + "/.env.toyotter")
	if err != nil {
		modules.ErrorMessage("Not found ~/.env.toyotter")
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

	name := `
	_                    _   _
 | |_ ___  _   _  ___ | |_| |_ ___ _ __
 | __/ _ \| | | |/ _ \| __| __/ _ \ '__|
 | || (_) | |_| | (_) | |_| ||  __/ |
	\__\___/ \__, |\___/ \__|\__\___|_|
					 |___/`

	tweetCommand := commands.TweetCommand(api, v)
	timelineCommand := commands.TimelineCommand(api, v)
	searchCommand := commands.SearchCommand(api, v)
	retweetCommand := commands.RetweetCommand(api, v)
	favoriteCommand := commands.FavoriteCommand(api, v)
	followCommand := commands.FollowCommand(api, v)
	blockCommand := commands.BlockCommand(api, v)
	mentionCommand := commands.MentionCommand(api, v)
	muteCommand := commands.MuteCommand(api, v)
	listCommand := commands.ListCommand(api, v)

	app := &cli.App{
		Name:    name,
		Usage:   "",
		Version: "0.5.6",
		Commands: []*cli.Command{
			&tweetCommand,
			&timelineCommand,
			&searchCommand,
			&retweetCommand,
			&favoriteCommand,
			&followCommand,
			&blockCommand,
			&mentionCommand,
			&muteCommand,
			&listCommand,
		},
	}

	args := os.Args

	if !terminal.IsTerminal(0) {
		b, _ := ioutil.ReadAll(os.Stdin)
		args = append(args, string(b))
	}

	err := app.Run(args)

	if err != nil {
		modules.ErrorMessage("Crash...")
	}
}
