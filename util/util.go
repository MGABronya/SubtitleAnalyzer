// @Title  util
// @Description  收集各种需要使用的工具函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// @title    ParseTime
// @description   尝试将line翻译为起始时间与终止时间
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     line		翻译行
// @return    start int, end int, err error		start为起始时间段，end为时间段，err为可能出现的错误
func ParseTime(line string) (start int, end int, err error) {
	// TODO 使用 strings.Split 函数，实现按 --> 分割字符串
	strArr := strings.Split(line, " --> ")
	if len(strArr) != 2 {
		err = errors.New("缺少符号' --> '，或者' --> '符号与前后时间的格式错误")
		return
	}
	start, err = Time(strArr[0])
	end, err = Time(strArr[1])
	// TODO 查看时间是否合理
	if start > end {
		err = errors.New("起始时间大于了终止时间")
	}
	return
}

// @title    Time
// @description   尝试将时间字符串翻译为时间
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     line		翻译行
// @return    time int, err error		time为时间字符串，err为可能出现的错误
func Time(line string) (time int, err error) {
	// TODO 使用 strings.Split 函数，实现按 : 分割字符串
	strArr := strings.Split(line, ":")
	if len(strArr) != 3 {
		err = errors.New("缺少符号':'或者':'附近符号格式出错")
		return
	}

	// TODO 查看小时格式
	if len(strArr[0]) != 2 {
		err = errors.New("小时格式错误")
		return
	}

	// TODO 读出小时
	hour, err := strconv.Atoi(strArr[0])

	if err != nil {
		err = errors.New("小时格式错误")
		return
	}
	// TODO 查看是否找出阈值
	if hour >= 60 {
		err = errors.New("小时超出阈值")
		return
	}

	// TODO 查看分钟格式
	if len(strArr[1]) != 2 {
		err = errors.New("分钟格式错误")
		return
	}

	// TODO 读出分钟
	minute, err := strconv.Atoi(strArr[1])
	if err != nil {
		err = errors.New("分钟格式错误")
		return
	}

	// TODO 查看是否找出阈值
	if minute >= 60 {
		err = errors.New("分钟超出阈值")
		return
	}

	// TODO 使用 strings.Split 函数，实现按 , 分割字符串
	strArr2 := strings.Split(strArr[2], ",")
	if len(strArr2) != 2 {
		err = errors.New("秒数与毫秒之间缺少','或者','前后格式错误")
		return
	}

	// TODO 查看秒格式
	if len(strArr2[0]) != 2 {
		err = errors.New("秒格式错误")
		return
	}

	// TODO 读出秒
	second, err := strconv.Atoi(strArr2[0])
	if err != nil {
		err = errors.New("秒格式错误")
		return
	}

	// TODO 查看是否找出阈值
	if second >= 60 {
		err = errors.New("秒超出阈值")
		return
	}

	// TODO 查看毫秒格式
	if len(strArr2[1]) != 3 {
		err = errors.New("毫秒格式错误")
		return
	}

	// TODO 读出毫秒
	msecond, err := strconv.Atoi(strArr2[1])
	if err != nil {
		err = errors.New("毫秒格式错误")
		return
	}

	// TODO 查看是否找出阈值
	if msecond >= 1000 {
		err = errors.New("毫秒超出阈值")
		return
	}

	// TODO 计算出时间
	time = hour*60*60*1000 + minute*60*1000 + second*1000 + msecond

	return
}

// @title    Time
// @description   尝试将时间翻译为字符串
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     time		需要翻译的时间
// @return    str string		str为时间字符串
func TimeToString(time int) (str string) {

	// TODO 计算毫秒
	ms := fmt.Sprint(time % 1000)
	// TODO 补零
	for len(ms) < 3 {
		ms = "0" + ms
	}
	time /= 1000

	// TODO 计算秒
	s := fmt.Sprint(time % 60)
	// TODO 补零
	for len(s) < 2 {
		s = "0" + s
	}
	time /= 60

	// TODO 计算分钟
	m := fmt.Sprint(time % 60)
	// TODO 补零
	for len(m) < 2 {
		m = "0" + m
	}
	time /= 60

	// TODO 计算小时
	h := fmt.Sprint(time)
	// TODO 补零
	for len(h) < 2 {
		h = "0" + h
	}

	str = h + ":" + m + ":" + s + "," + ms
	return
}
