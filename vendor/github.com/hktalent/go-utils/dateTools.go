package go_utils

import (
	"fmt"
	"time"
)

// 利用各种格式自动解析
func Parse2Time(s string) *time.Time {
	var aLy = []string{time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly}

	// 使用time.ParseInLocation()函数解析时间字符串
	for _, timeStr := range aLy {
		t, err := time.ParseInLocation(timeStr, s, time.Local)
		if err != nil {
			fmt.Println(err)
		} else {
			return &t
		}
	}
	return nil
}
