env:
	go run main.go env

build:
        go build -o ktools main.go
        sudo chmod +rwx ./ktools
        sudo cp ./ktools /bin/

