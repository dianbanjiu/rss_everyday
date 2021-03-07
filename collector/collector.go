package collector

import (
	"RSS_bot/cmd"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"
	"os"
	"time"
)

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
	now := time.Now()
	yesterday := time.Date(now.Year(),now.Month(), now.Day()-1, 12, 01, 00, 00, now.Location()).Unix()
	today := time.Date(now.Year(), now.Month(), now.Day(), 12, 00, 59, 00, now.Location()).Unix()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rss.Url)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		for _, item := range feed.Items {
			if item.PublishedParsed.Unix() >= yesterday && item.PublishedParsed.Unix() <= today {
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
	bot, err := tgbotapi.NewBotAPI(*cmd.BotToken)
	if err != nil {
		panic(err)
	}
	for _, s := range msg {
		_, _ = bot.Send(tgbotapi.NewMessage(*cmd.ChannelID, s))
	}

}
