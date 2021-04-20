# 进击的 RSS
RSS_EVERYDAY 是一个 RSS 订阅工具，它会每四小时采集一次新的文章，并通过 TG BOT 推送至指定的频道/群组。
### [中文说明](https://github.com/GuangzheJiang/rss_everyday/blob/main/README.md)|[English Description](https://github.com/GuangzheJiang/rss_everyday/blob/main/README_en.md)
## How to use
## 项目使用姿势

### TG Bot 创建
添加 `@BotFather` 机器人，注册一个 bot，记录下 bot 的 token  

### TG Channel 创建
创建一个 Channel，添加之前创建的 bot 为管理员，添加方式：  
点击频道右上角的`管理频道 - 管理员 - 添加管理员`，搜索你创建的 bot 名称，将其添加为管理员即可，发送消息的权限必须赋予，其他权限可选。  

[点此登陆](https://web.telegram.org) 网页版 telegram。点击你刚才创建的频道，链接格式类似 `https://web.telegram.org/#/im?p=cxxx_ppp` ，在这种情况下 `-100xxx` 为你的 Channel ID。  

### TG Group 创建
创建一个群组，添加之前创建的 bot 为管理员，添加方式类似 Channel。  

[点此登陆](https://web.telegram.org) 网页版 telegram。点击你刚才创建的群组，链接格式类似`https://web.telegram.org/#/im?p=sxxx_ppp` ，在这种情况下 `-100xxx` 为你的 Group ID。
或者也有可能类似 `https://web.telegram.org/#/im?p=gxxx` ，在这种情况下 `-xxx` 为你的 Group ID。  

### github 配置
克隆本仓库，在仓库的 `Settings-Secrets` 新增两个字段 `BOTTOKEN` 和 `CHANNELID` ，他们的值分别是你的 bot token 和 Channel ID/Group ID。


### RSS 源添加

如果网站本身提供了 RSS 链接，你可以按照 rss.json 的格式直接添加。如果没有提供 RSS 链接，你可以通过 [RSSHub](https://docs.rsshub.app/) 制作对应的链接。

## 注意事项
代码中是使用发布时间进行筛选的，如果你所订阅的 RSS 链接无法解析出该字段，建议不要将该链接添加至 rss.json 文件。

此项目是基于 Github Action 的定时任务实现的，并且只会在达到设定时间的情况下才会触发 Action ，也可以手动触发。
