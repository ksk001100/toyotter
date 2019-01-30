# toyotter2
![](https://img.shields.io/github/stars/KeisukeToyota/toyotter2.svg)
![](https://img.shields.io/github/release/KeisukeToyota/toyotter2.svg)
![](https://img.shields.io/github/issues/KeisukeToyota/toyotter2.svg)
![](https://img.shields.io/github/forks/KeisukeToyota/toyotter2.svg)
![](https://img.shields.io/github/license/KeisukeToyota/toyotter2.svg)
[![Build Status](https://travis-ci.org/KeisukeToyota/toyotter2.svg?branch=master)](https://travis-ci.org/KeisukeToyota/toyotter2)
[![Go Report Card](https://goreportcard.com/badge/github.com/KeisukeToyota/toyotter2)](https://goreportcard.com/report/github.com/KeisukeToyota/toyotter2)

CUI based Twitter client made with Golang

![Imgur](https://i.imgur.com/JjzlCnh.png)

## Install

```shell
$ go get github.com/KeisukeToyota/toyotter2
$ cp $GOPATH/src/github.com/KeisukeToyota/toyotter2/.env.example ~/.env.toyotter
```

In macOS you can install it with Homebrew
```shell
$ brew tap keisuketoyota/homebrew-toyotter2
$ brew install toyotter2
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
$ toyotter2 tweet "Hello world" #=> Tweet "Hello world"
$ toyotter2 tw "Bye" #=> Tweet "Bye"
$ toyotter2 tweet "Hello world" --image image.jpg #=> Tweet image
$ toyotter2 tw "Hello" -img=image.jpg #=> Tweet image
$ toyotter2 tweet delete 34820348023 #=> Delete the 34820348023 tweet
$ toyotter2 tw d 34820348023 #=> Delete the 34820348023 tweet
$ toyotter2 tweet "Hi" --reply 34820348023 #=> Reply to tweet of 34820348023
$ toyotter2 tw "Bye" --reply 34820348023 #=> Reply to tweet of 34820348023
```

### Timeline
```shell
$ toyotter2 timeline #=> Get timeline(Get 5 by default)
$ toyotter2 timeline --count 30 #=> Get 30 timelines
$ toyotter2 tl -c=20 #=> Get 20 timelines
```

### Search
```shell
$ toyotter2 search --user twitter #=> Get users including "twitter"(Get 5 by default)
$ toyotter2 s --tweet twitter #=> Get users including "twitter"(Get 5 by default)
$ toyotter2 s -u twitter --count 20 #=> Get 20 users including twitter
$ toyotter2 search -tw=twitter -c=30 #=> Get 30 users including twitter
```

### Retweet
```shell
$ toyotter2 retweet 34820348023 #=> Retweet tweet of ID 34820348023
$ toyotter2 rt 34820348023 #=> Retweet tweet of ID 34820348023
$ toyotter2 retweet 34820348023 --delete #=> UnRetweet tweet of ID 34820348023
$ toyotter2 rt 34820348023 -d #=> UnRetweet tweet of ID 34820348023
```

### Favorite
```shell
$ toyotter2 favorite 34820348023 #=> Favorite tweet of ID 34820348023
$ toyotter2 fav 34820348023 #=> Favorite tweet of ID 34820348023
$ toyotter2 favorite 34820348023 --delete #=> UnFavorite tweet of ID 34820348023
$ toyotter2 fav 34820348023 -d #=> UnFavorite tweet of ID 34820348023
```

### Follow
```shell
$ toyotter2 follow TwitterJP #=> Follow @TwitterJP
$ toyotter2 flw TwitterJP #=> Follow @TwitterJP
$ toyotter2 follow TwitterJP --delete #=> UnFollow @TwitterJP
$ toyotter2 flw TwitterJP -d #=> UnFollow @TwitterJP
```

### Block
```shell
$ toyotter2 block TwitterJP #=> Block @TwitterJP
$ toyotter2 blk TwitterJP #=> Block @TwitterJP
$ toyotter2 block TwitterJP --delete #=> UnBlock @TwitterJP
$ toyotter2 blk TwitterJP -d #=> UnBlock @TwitterJP
$ toyotter2 block --list #=> Get block user list
```

### Mention list
```shell
$ toyotter2 mention #=> Get mentions(Get 5 by default)
$ toyotter2 men #=> Get mentions(Get 5 by default)
$ toyotter2 mention --count 20 #=> Get 20 mentions
$ toyotter2 men -c 20 #=> Get 20 mentions
```

### Mute
```shell
$ toyotter2 mute TwitterJP #=> Mute @TwitterJP
$ toyotter2 mu TwitterJP #=> Mute @TwitterJP
$ toyotter2 mute TwitterJP --delete #=> UnMute @TwitterJP
$ toyotter2 mu TwitterJP -d #=> UnMute @TwitterJP
```

### Help
```shell
$ toyotter2 --help #=> Overall help
$ toyotter2 -h #=> Overall help
$ toyotter2 timeline --help #=> Timeline command help
$ toyotter2 rt -h #=> Retweet command help
```
