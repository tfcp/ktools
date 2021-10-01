env:
	go run main.go env

build:
        go build -o ktools main.go

init:
    ln -s /Users/zhaosuji/go/src/github.com/ktools/ktools /usr/local/bin/ktools
