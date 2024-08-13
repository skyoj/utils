package utils

import (
	"encoding/json"
	"github.com/microcosm-cc/bluemonday"
	"strconv"
)

// S 字符串类型转换
type S string

func (s S) String() string {
	return string(s)
}

// Bytes 转换为[]byte
func (s S) Bytes() []byte {
	return []byte(s)
}

// Int64 转换为int64
func (s S) Int64() int64 {
	i, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Int 转换为int
func (s S) Int() int {
	return int(s.Int64())
}

// Uint 转换为uint
func (s S) Uint() uint {
	return uint(s.Uint64())
}

// Uint 转换为uint
func (s S) Uint32() uint32 {
	return uint32(s.Uint64())
}

// Uint64 转换为uint64
func (s S) Uint64() uint64 {
	i, err := strconv.ParseUint(s.String(), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Float64 转换为float64
func (s S) Float64() float64 {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		return 0
	}
	return f
}

// ToJSON 转换为JSON
func (s S) ToJSON(v interface{}) error {
	return json.Unmarshal(s.Bytes(), v)
}

// MarkdownToHTML 将markdown 转换为 html
//func (s S) MarkdownToHTML() string {
//	unsafebytes := blackfriday.Run([]byte(s.String()))
//	return string(bluemonday.UGCPolicy().SanitizeBytes(unsafebytes))
//}

// 避免XSS
func (s S) AvoidXSS() string {
	p := bluemonday.UGCPolicy()
	p.AllowImages()
	p.AllowDataURIImages()
	return p.Sanitize(s.String())
}

// 去除重复字符串
func RemoveRepeatedElement(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// ReplaceRightSingle
// 从右往左移除指定字符(单个字符)
// str 原字符串
// substr 移除字符串(单个字符, 可指定多个单字符), 多个时只移除一个; 单字符数组: old 不支持 []string{"abc"}, 支持[]string{"a", "b", "c"}
// n 移除数量
func ReplaceRightSingle(str string, n int, new string, old []string) string {
	if n == 0 || len(old) == 0 {
		return str
	}

	var count int
	var newStr = str
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		for _, sub := range old {
			if string(runes[i]) == sub {
				newStr = string([]rune(newStr)[:i]) + new + string([]rune(newStr)[i+1:])
				count++
			}
			if count == n {
				return newStr
			}
		}
	}

	return newStr
}
