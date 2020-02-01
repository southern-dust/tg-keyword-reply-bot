# telegram 关键词自动回复机器人 开源版本

这个是关键词回复机器人的开源版本代码，感谢原作者 @zu1k 的开发。在开源代码基础上维护和更新了一些代码 
## 开源版本介绍
### 基本命令介绍
- 添加关键词回复规则 `/add 关键词===回复内容` 或者 `/add 关键词1||关键词2===回复内容` 
- 关键词可以使用正则表达式,例如`/add re:p([a-z]+)ch===测试正则`,就会匹配规则`p([a-z]+)ch` 
- 关键词可以使用同一目录下的 rules.txt 添加，格式为 `<groupid> rule`，一行为一条规则（或多条（使用正则表达式））
- 删除关键词规则 `/del 关键词` 暂不支持一次性删除多个关键词
- 自动删除含有关键词的文字消息, 只需要将回复内容设置成 `delete`, 并给机器人添加删除消息权限
- 使用`/list`命令可以查看本群内所有自动回复规则
- 给机器人添加删除消息和踢人的管理权限,可以自动防清真(阿拉伯语)

### 使用开源版本编译搭建
- 建议使用最近的 go 版本进行编译
- 初次使用 `./tg-keyword-reply <tg-bot-token>` , 会将token存到数据库中(bot.db)
- 后面使用 `./tg-keyword-reply` , 无需输入token

### 回复特殊内容
- 回复内容支持文字\图片\GIF\视频,默认文字
- 如需图片,回复内容设置成`photo:https://t.me/c/1472018167/53095`,`https://t.me/c/1472018167/53095`是已经发送过的图片获取到的链接
- 同理,gif将`photo`替换成`gif`,视频替换成`video`,文件替换成`file`
- 注意: 这里的链接必须是公开群组的,否则无法发出来



#### 机器人命令列表
是推荐给用户显示的命令，私聊机器人爹设置

```
help - 查看帮助
add - 添加规则
del - 删除规则
list - 列出规则
admin - 呼叫管理员
banme - 禁言小游戏
getid - 查看用户的信息 可回复查看别人
autoreply - 开关自动回复功能  //TODO
autodelete - 开关自动删除消息功能  //TODO
replyorder - 开关回复ban/kick命令功能  //TODO
banmegame - 开关禁言小游戏功能  //TODO
playorderban - 开关玩命令惩罚功能  //TODO
banqingzhen - 开关防清真功能  //TODO
calladmin - 开关呼叫管理员功能  //TODO
welcome - 开关加群欢迎功能  //TODO
goodbye - 开关离群送别功能  //TODO
deletejoinmessage - 开关删除加群消息功能  //TODO
servicelist - 查看机器人功能列表  //TODO
```
