package tool

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func RandomKey(t int64) string {
	rand.Seed(t)
	md5str := string([]byte(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", rand.Int31())))))[:8])
	return md5str
}

// 获取起始日期内所有日期
func GetBetweenDates(startDate, endDate string, layout string) []string {
	d := []string{}
	if startDate == endDate { // 这个地方有问题
		d = append(d, startDate)
		return d
	}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = layout
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// Decimal 浮点数保留小数
func Decimal(value float64, p int) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(p)+"f", value), 64)
	return value
}

func Money(value float64) string {
	strnum := strconv.FormatFloat(value, 'f', 2, 64)
	capitalSlice := []string{"万", "亿", "兆"}
	index := 0
	result := ""
	sdivision := strings.Split(strnum, ".")
	sl := sdivision[0]
	if len(sdivision) > 1 {
		result = "." + sdivision[1]
	}
	// slength := len(sl)
	for len(sl) > 4 {
		result = capitalSlice[index] + sl[len(sl)-4:] + result
		index = index + 1
		sl = sl[0 : len(sl)-4]
	}
	result = sl + result
	result = strings.Replace(result, "万0000", "万", -1)
	result = strings.Replace(result, "亿0000", "亿", -1)
	result = strings.Replace(result, "兆0000", "兆", -1)
	result = strings.Replace(result, "亿万", "亿", -1)
	result = strings.Replace(result, "兆亿", "兆", -1)
	if strings.Contains(result, "万") {
		strs := strings.Split(result, "万")
		result = strs[0] + "万"
	}
	return result
}

// 获取模型字段
func GetModelFields(i interface{}) string {
	obj := reflect.TypeOf(i)
	fields := make([]string, 0)
	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		fields = append(fields, field.Tag.Get("json"))
	}
	fieldStr := strings.Join(fields, ",")
	return fieldStr
}
