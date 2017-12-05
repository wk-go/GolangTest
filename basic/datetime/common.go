package main

/**
time常用方法

After(u Time) bool
时间类型比较,是否在Time之后

Before(u Time) bool
时间类型比较,是否在Time之前

Equal(u Time) bool
比较两个时间是否相等

IsZero() bool
判断时间是否为零值,如果sec和nsec两个属性都是0的话,则该时间类型为0

Date() (year int, month Month, day int)
返回年月日,三个参数

Year() int
返回年份

Month() Month
返回月份.是Month类型

Day() int
返回多少号

Weekday() Weekday
返回星期几,是Weekday类型

ISOWeek() (year, week int)
返回年份,和该填是在这年的第几周.

Clock() (hour, min, sec int)
返回小时,分钟,秒

Hour() int
返回小时

Minute() int
返回分钟

Second() int
返回秒数

Nanosecond() int
返回纳秒
 */

import "fmt"
import "time"

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2017, 06, 21, 20, 34, 58, 0, time.UTC)
    p(then)

    p("Year:",then.Year())
    p("Month:",int(then.Month()))
    p("Day:",then.Day())
    p("Hour:",then.Hour())
    p("Minute:",then.Minute())
    p("Second:",then.Second())
    p("Nanosecond:",then.Nanosecond())
    p("Location:",then.Location())

    p("Weekday:",then.Weekday())

    p("Before:",then.Before(now))
    p("After:",then.After(now))
    p("Equal:",then.Equal(now))

    diff := now.Sub(then)
    p("diff:",diff)

    p("diff.Hours:",diff.Hours())
    p("diff.Minutes:",diff.Minutes())
    p("diff.Seconds:",diff.Seconds())
    p("diff.Nanoseconds:",diff.Nanoseconds())

    p("then.Add+diff:",then.Add(diff))
    p("then.Add-diff:",then.Add(-diff))
}