## What's this project

To invite email to each 3rd party by controlling a web browser or using api of 3rd party.

## Step by step

```
brew install chromedriver \
             selenium-server-standalone \
             direnv
```

```
go mod init
go get github.com/sclevine/agouti
```

```
mv .envrc.sample .envrc
direnv allow .
```

## For development

```
docker-compse up -d

docker-compose logs -f
```

## go run main.go

```
docker-compose run --rm slack-bot go run *.go
```
