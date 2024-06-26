package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now()
	now := time.Now()
	fmt.Println(now)

	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	fmt.Printf("%d年%d月%d日 %02d:%02d:%02d\n", year, month, day, hour, min, sec)

	// 時間の加算や減算
	oneHourLater := now.Add(time.Hour)
	fmt.Println(oneHourLater)

	// 2つの時間の差
	start := time.Now()
	// 何らかの処理（今回は1秒スリープ）
	time.Sleep(1 * time.Second) // 1秒スリープ
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("処理にかかった時間:", duration)

	formatted := now.Format("2006年01月02日 15時04分05秒")
	fmt.Println(formatted)

	timeStr := "2023-04-02 15:30:00"
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("時間のパースに失敗しました:", err)
		return
	}
	fmt.Println(t)
}
