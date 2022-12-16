// @Title  subtitle
// @Description  定义字幕
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

// Subtitle			定义字幕
type Subtitle struct {
	Id        int    `json:"id"`         // 用户Id
	StartTime int    `json:"start_time"` // 起始时间
	EndTime   int    `json:"end_time"`   // 终止时间
	Content   string `json:"content"`    // 内容
}

// TODO 定义排序方式
type Subtitles []Subtitle

func (s Subtitles) Len() int { return len(s) }

func (s Subtitles) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Subtitles) Less(i, j int) bool { return s[i].StartTime < s[j].StartTime }
