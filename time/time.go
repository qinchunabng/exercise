package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t) //一秒钟执行一次
	// }
	//时间格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
