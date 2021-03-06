package main

import (
	"RSS_bot/cmd"
	"RSS_bot/collector"
	"flag"
)

func init() {
	cmd.BotToken = flag.String("tg_bot", "", "Telegram bot token")
	cmd.ChannelID = flag.Int64("tg_channel", 0, "Telegram channel id")
	flag.Parse()
	cmd.TokenVaild()
}

func main() {


	collector.GetPosts()
}
