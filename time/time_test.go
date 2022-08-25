package time

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTimeParse2Day(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Format("20060102150405"))
}

func Test(t *testing.T) {
	//获取本地location
	toBeCharge := "10-01-2017"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "01-02-2006"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	// 时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)

	fmt.Println(time.Now().Add(-1).Before(time.Now()))

	// -62135596800,
	newTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", strconv.FormatInt( 1583045182,10))
	fmt.Println(newTime)

	fmt.Println(strconv.FormatInt( 1583045182,10))

	fmt.Println(time.Unix(1583045182, 0).Format("2006-01-02T15:04:05Z07:00"))

	////////////////////////////////////////////////////////////////////////////////////////////////////
}

func BenchmarkLocation(b *testing.B) {
    for n := 0; n < b.N; n++ {
        loc, _ := time.LoadLocation("Asia/Kolkata")
        time.Now().In(loc)
    }
}

func BenchmarkLocation2(b *testing.B) {
    loc, _ := time.LoadLocation("Asia/Kolkata")
    for n := 0; n < b.N; n++ {
        time.Now().In(loc)
    }
}