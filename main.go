package main

import (
	"net/url"
	"os"
	"os/user"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/commands"
	"github.com/KeisukeToyota/toyotter/modules"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
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

	app := cli.NewApp()

	app.Name = `
	_                    _   _
 | |_ ___  _   _  ___ | |_| |_ ___ _ __
 | __/ _ \| | | |/ _ \| __| __/ _ \ '__|
 | || (_) | |_| | (_) | |_| ||  __/ |
	\__\___/ \__, |\___/ \__|\__\___|_|
					 |___/`

	app.Usage = ""
	app.Version = "0.3.6"
	app.Commands = []cli.Command{
		commands.TweetCommand(api, v),
		commands.TimelineCommand(api, v),
		commands.SearchCommand(api, v),
		commands.RetweetCommand(api, v),
		commands.FavoriteCommand(api, v),
		commands.FollowCommand(api, v),
		commands.BlockCommand(api, v),
		commands.MentionCommand(api, v),
		commands.MuteCommand(api, v),
		commands.ListCommand(api, v),
	}

	err := app.Run(os.Args)
	if err != nil {
		modules.ErrorMessage("Crash...")
	}
}
