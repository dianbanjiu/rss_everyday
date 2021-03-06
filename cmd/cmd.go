package cmd

var BotToken *string
var ChannelID *int64

func TokenVaild() {
	if *BotToken == "" || *ChannelID == 0 {
		panic("BotToken && ChannelID cannot be empty")
	}
}