package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderId  uint   //发送者
	ReceverId uint   //接收者
	Type      int    //发送类型 群聊 私聊 广播
	Media     string //消息类型 文字 图片 音频
	Content   string //消息内容
	Picture   string
	Url       string
	Desc      string
	Amount    int
}

func (table *Message) TableName() string {
	return "message"
}
