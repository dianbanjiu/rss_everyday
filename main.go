package main

import (
	"encoding/json"
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"
	"os"
	"time"
)

// 基础环境配置
var BotToken *string
var ChannelID *int64

func TokenValid() {
	if *BotToken == "" || *ChannelID == 0 {
		panic("BotToken && ChannelID cannot be empty")
	}
}

func init() {
	BotToken = flag.String("tg_bot", "", "Telegram bot token")
	ChannelID = flag.Int64("tg_channel", 0, "Telegram channel id")
	flag.Parse()
	TokenValid()
	GetRssInfo()
}

// RSS 构成阶段
type RSSInfos struct {
	RssInfo []RssInfo `json:"rss_info"`
}
type RssInfo struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	FullContent bool   `json:"full_content"`
}

var RssInfos = RSSInfos{nil}

// 从 配置文件中获取 rss 链接
// 根据 rss 链接获取更新
func GetRssInfo() {
	rssFile, err := os.Open("./rss.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(rssFile).Decode(&RssInfos)
	if err != nil {
		panic(err)
	}

}

// 根据时间筛选昨天一整天的文章
func GetPosts() {
	var msg = make([]string, 0)
	for _, info := range RssInfos.RssInfo {
		msg = append(msg, GetPostInfo(info)...)
	}
	PushPost(msg)
}

func GetPostInfo(rss RssInfo) []string {
	var msg = make([]string, 0)
	now := time.Now().UTC()
	startTime := now.Add(-4 * time.Hour)
	start := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), 0, 0, 0, now.Location()).Unix()
	end := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).Unix()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rss.Url)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		for _, item := range feed.Items {
			if item.PublishedParsed != nil && item.PublishedParsed.Unix() >= start && item.PublishedParsed.Unix() < end {
				msgItem := fmt.Sprintln(item.Title, item.Link)
				msg = append(msg, msgItem)

			}
		}
	}

	return msg
}

// 从配置文件获取推送方式
// 使用对应的推送渠道推送文章
func PushPost(msg []string) {
	bot, err := tgbotapi.NewBotAPI(*BotToken)
	if err != nil {
		panic(err)
	}
	for _, s := range msg {
		_, _ = bot.Send(tgbotapi.NewMessage(*ChannelID, s))
	}

}

func main() {
	GetPosts()
}
