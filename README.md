# toyotter
![](https://img.shields.io/github/stars/ksk001100/toyotter.svg)
![](https://img.shields.io/github/release/ksk001100/toyotter.svg)
![](https://img.shields.io/github/issues/ksk001100/toyotter.svg)
![](https://img.shields.io/github/forks/ksk001100/toyotter.svg)
![](https://img.shields.io/github/license/ksk001100/toyotter.svg)
[![Build Status](https://travis-ci.org/ksk001100/toyotter.svg?branch=master)](https://travis-ci.org/ksk001100/toyotter)
[![Go Report Card](https://goreportcard.com/badge/github.com/ksk001100/toyotter)](https://goreportcard.com/report/github.com/ksk001100/toyotter)

CUI based Twitter client made with Golang

![Imgur](https://i.imgur.com/4iLQ5Bu.png)

## Install

```shell
$ go get github.com/ksk001100/toyotter
$ cp $GOPATH/src/github.com/ksk001100/toyotter/.env.example ~/.env.toyotter
```

In macOS you can install it with Homebrew
```shell
$ brew tap ksk001100/toyotter
$ brew install toyotter
```

Please set up Twitter key to `~/.env.toyotter`
```
CONSUMER_KEY=xxxxxxxxxxxxxxxxxxxxxxx
CONSUMER_SECRET=xxxxxxxxxxxxxxxxxxxxxxx
ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxx
ACCESS_TOKEN_SECRET=xxxxxxxxxxxxxxxxxxxxxxx
```

## Usage

### Tweet
```shell
$ toyotter tweet "Hello world" #=> Tweet "Hello world"
$ toyotter tw "Bye" #=> Tweet "Bye"
$ toyotter tweet "Hello world" --image image.jpg #=> Tweet image
$ toyotter tw "Hello" -img=image.jpg #=> Tweet image
$ toyotter tweet delete 34820348023 #=> Delete the 34820348023 tweet
$ toyotter tw d 34820348023 #=> Delete the 34820348023 tweet
$ toyotter tweet "Hi" --reply 34820348023 #=> Reply to tweet of 34820348023
$ toyotter tw "Bye" --reply 34820348023 #=> Reply to tweet of 34820348023
```

### Timeline
```shell
$ toyotter timeline #=> Get timeline(Get 5 by default)
$ toyotter timeline --count 30 #=> Get 30 timelines
$ toyotter tl -c=20 #=> Get 20 timelines
```

### Search
```shell
$ toyotter search --user twitter #=> Get users including "twitter"(Get 5 by default)
$ toyotter s --tweet twitter #=> Get users including "twitter"(Get 5 by default)
$ toyotter s -u twitter --count 20 #=> Get 20 users including twitter
$ toyotter search -tw=twitter -c=30 #=> Get 30 users including twitter
```

### Retweet
```shell
$ toyotter retweet 34820348023 #=> Retweet tweet of ID 34820348023
$ toyotter rt 34820348023 #=> Retweet tweet of ID 34820348023
$ toyotter retweet 34820348023 --delete #=> UnRetweet tweet of ID 34820348023
$ toyotter rt 34820348023 -d #=> UnRetweet tweet of ID 34820348023
```

### Favorite
```shell
$ toyotter favorite 34820348023 #=> Favorite tweet of ID 34820348023
$ toyotter fav 34820348023 #=> Favorite tweet of ID 34820348023
$ toyotter favorite 34820348023 --delete #=> UnFavorite tweet of ID 34820348023
$ toyotter fav 34820348023 -d #=> UnFavorite tweet of ID 34820348023
```

### Follow
```shell
$ toyotter follow TwitterJP #=> Follow @TwitterJP
$ toyotter flw TwitterJP #=> Follow @TwitterJP
$ toyotter follow TwitterJP --delete #=> UnFollow @TwitterJP
$ toyotter flw TwitterJP -d #=> UnFollow @TwitterJP
```

### Block
```shell
$ toyotter block TwitterJP #=> Block @TwitterJP
$ toyotter blk TwitterJP #=> Block @TwitterJP
$ toyotter block TwitterJP --delete #=> UnBlock @TwitterJP
$ toyotter blk TwitterJP -d #=> UnBlock @TwitterJP
$ toyotter block --list #=> Get block user list
```

### Mention list
```shell
$ toyotter mention #=> Get mentions(Get 5 by default)
$ toyotter men #=> Get mentions(Get 5 by default)
$ toyotter mention --count 20 #=> Get 20 mentions
$ toyotter men -c 20 #=> Get 20 mentions
```

### Mute
```shell
$ toyotter mute TwitterJP #=> Mute @TwitterJP
$ toyotter mu TwitterJP #=> Mute @TwitterJP
$ toyotter mute TwitterJP --delete #=> UnMute @TwitterJP
$ toyotter mu TwitterJP -d #=> UnMute @TwitterJP
```

### Help
```shell
$ toyotter --help #=> Overall help
$ toyotter -h #=> Overall help
$ toyotter timeline --help #=> Timeline command help
$ toyotter rt -h #=> Retweet command help
```
