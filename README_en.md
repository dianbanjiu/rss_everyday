RSS_EVERYDAY is an RSS subscription tool. It collects new articles every four hours and pushes them to designated channels/group via TG BOT.
### [中文说明](https://github.com/GuangzheJiang/rss_everyday/blob/main/README.md)|[English Description](https://github.com/GuangzheJiang/rss_everyday/blob/main/README_en.md)
## How to use

### Create TG BOT
Add the `@BotFather` robot, register a bot, and remember the token of the bot.

### Create TG Channel
Create a Channel, add the previously created bot as an administrator:

> Click `Manage Channel-Administrator-Add Administrator` in the upper right corner of the channel, search for the name of the bot you created, and add it as an administrator. The permission to send messages must be granted, and other permissions are optional.

[Login here](https://web.telegram.org) The web version of telegram. Click on the channel you just created. The link format is similar to `https://web.telegram.org/#/im?p=cxxx_ppp`, where `-100xxx` is your channel ID, write it down, and use it later.

### Create TG Group
Create a group. The way add bot to group is same with channel. 

[Login here](https://web.telegram.org) The web version of telegram. Click on the channel you just created. The link format is similar to `https://web.telegram.org/#/im?p=gxxx`, where `-xxx` is your group ID. Or maybe it is similar to `https://web.telegram.org/#/im?p=sxxx_ppp`, in which case `-100xxx` is your group ID.

### Github Configuration
Clone this repository and add two new fields `BOTTOKEN` and `CHANNELID` to the `Settings-Secrets` of the repository. Their values are your bot token and channel/group ID respectively.

### RSS feed addition

If the website itself provides an RSS link, you can directly add it in the format of rss.json. If no RSS link is provided, you can make the corresponding link through [RSSHub](https://docs.rsshub.app/).

## Precautions
The code uses the publication time to filter. If the RSS link you subscribe to cannot parse the field, it is recommended not to add the link to the rss.json file.

This project is based on the timing task of Github Action, and the Action will only be triggered when the set time is reached. You can also manually trigger the action to test its availability.
