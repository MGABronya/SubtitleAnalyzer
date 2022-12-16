// @Title  SubtitleController
// @Description  该文件提供关于操作字幕的各种方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"SubtitleAnalyzer/model"
	"SubtitleAnalyzer/response"
	"SubtitleAnalyzer/util"
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @title    Upload
// @description   上传待解析的字幕文件
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func Upload(ctx *gin.Context) {

	// TODO 读取文件
	file, err := ctx.FormFile("file")

	//TODO 数据验证
	if err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 打开文件
	f, err := file.Open()
	f.Close()

	//TODO 数据验证
	if err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 按行读入文件
	buff := bufio.NewReader(f)

	// TODO 定义字幕数组
	subtitles := make([]model.Subtitle, 0)

	// TODO 定义单个字幕
	subtitle := model.Subtitle{}

	// TODO 定义报错
	errors := make([]string, 0)

	// TODO 定义警报
	warnings := make([]string, 0)

	// TODO 定义状态
	state := 0

	// TODO 解析字幕
	for i := 1; ; i++ {
		ln, _, err := buff.ReadLine()
		len := len(ln) - 1
		for ; len >= 0; len-- {
			if ln[len] != '\r' && ln[len] != '\n' && ln[len] != ' ' {
				break
			}
		}
		line := string(ln[0 : len+1])
		// TODO 读到了文件尾
		if io.EOF == err {
			// TODO 缺失内容
			if state == 1 {
				warnings = append(warnings, fmt.Sprintf("文件在第%d行格式有误，文件尾缺失内容", i))
				break
			}
			break
		}
		// TODO 查看是否读出失败
		if err != nil {
			errors = append(errors, fmt.Sprintf("文件在第%d行格式有误，读出失败", i))
			continue
		} else {
			if state == 0 {
				// TODO 此处允许多余的空行
				if line == "" {
					warnings = append(warnings, fmt.Sprintf("文件在第%d处存在多余的空行", i))
					continue
				}
				id, err := strconv.Atoi(line)

				// TODO id格式出错
				if err != nil || id < 0 {
					errors = append(errors, fmt.Sprintf("文件在第%d处id格式出错，此处应为大于0的整数", i))
					// TODO 标记为损坏
					id = -1
				}
				// TODO 置入subtitle中
				subtitle.Id = id
				// TODO 控制状态转移
				state++
			} else if state == 1 {

				// TODO 尝试读出时间
				start, end, err := util.ParseTime(line)
				// TODO 格式出错
				if err != nil {
					errors = append(errors, fmt.Sprintf("文件在第%d行时间格式有误，"+err.Error(), i))
					// TODO 标记为损坏
					subtitle.Id = -1
				}
				// TODO 置入subtitle中
				subtitle.StartTime = start
				subtitle.EndTime = end
				// TODO 控制状态转移
				state++
			} else {
				// TODO 读入行
				// TODO 空行时终止读入
				if line == "" {
					// TODO 为损坏数据，则不压入数组
					if subtitle.Id < 0 {
						continue
					}
					// TODO 将subtitle压入数组
					subtitles = append(subtitles, subtitle)
					// TODO 清空subtitle
					subtitle = model.Subtitle{}
					// TODO 控制状态转移
					state = 0
					continue
				}
				subtitle.Content = subtitle.Content + line + "\n"
			}
		}
	}

	set := make(map[int]bool, 0)
	// TODO 对subtitles进行编号检查
	for i := 0; i < len(subtitles); i++ {
		if set[subtitles[i].Id] {
			warnings = append(warnings, fmt.Sprintf("字幕编号%d出现冲突", subtitles[i].Id))
		}
		set[subtitles[i].Id] = true
	}

	// TODO 对subtitles进行排列
	sort.Sort(model.Subtitles(subtitles))

	// TODO 检查subtitles的时间区间是否出现重叠
	for i := 1; i < len(subtitles); i++ {
		if subtitles[i].StartTime < subtitles[i-1].EndTime {
			warnings = append(warnings, fmt.Sprintf("字幕编号%d与字幕id%d的时间区间出现冲突", subtitles[i-1].Id, subtitles[i].Id))
			// TODO 此处压缩重叠的字幕时间，使字幕错开
			subtitles[i].StartTime = subtitles[i-1].EndTime
			if subtitles[i].EndTime < subtitles[i-1].EndTime {
				subtitles[i].EndTime = subtitles[i-1].EndTime
			}
		}
	}

	// TODO 成功
	response.Success(ctx, gin.H{"subtitles": subtitles, "errors": errors, "warnings": warnings}, "分析字幕成功")
}
