package wechatrequests

type MarkdownMessage struct {
	Markdown struct {
		Content string `json:"content" gorm:"column:content"`
	} `json:"markdown" gorm:"column:markdown"`
	Msgtype string `json:"msgtype" gorm:"column:msgtype"`
}
