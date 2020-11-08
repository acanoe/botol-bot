# botol-bot
A telegram bot written in Go, using Cobra and Telebot

## How to use
You can build the executable first by using this command:
```
go build
```
or directly run it:
```
go run main.go
```
To start the bot pass your bot token to the `--botToken` (or `-k`) flag into the `start` command, e.g.:
```
./botol-bot start --botToken <your bot token> #or directly without building it:
go run main.go start --botToken <your bot token>
```
