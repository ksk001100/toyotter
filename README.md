# toyotter2

Golang製のCUIベースTwitterクライアント

## インストール

```
$ go get github.com/KeisukeToyota/toyotter2
$ cp $GOPATH/src/github.com/KeisukeToyota/toyotter2/.env.example ~/.env.toyotter 
```
macOSはHomebrewでもインストールできます
```
$ brew tap keisuketoyota/homebrew-toyotter2
$ brew install toyotter2
```

`~/.env.toyotter` にTwitterのキーを設定してください
```
CONSUMER_KEY=xxxxxxxxxxxxxxxxxxxxxxx
CONSUMER_SECRET=xxxxxxxxxxxxxxxxxxxxxxx
ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxx
ACCESS_TOKEN_SECRET=xxxxxxxxxxxxxxxxxxxxxxx
```

## 使い方
`$ toyotter2 --help` を参照してください
