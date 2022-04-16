package main

import (
	"fmt"
	"time"
	"week05/pkg/window"
)

func main() {
	win := window.NewWindow(window.Size(10), window.Duration(time.Millisecond*100))

	//每100ms的请求数值落到下一个桶
	for i := 0; i < 20; i++ {
		win.Add(i + 1)
		fmt.Println(win)
		<-time.After(time.Millisecond * 100)
	}

	//每100ms统计sum会减少一个桶的数值
	for i := 0; i < 20; i++ {
		fmt.Println("sum", win.Sum())
		<-time.After(time.Millisecond * 100)
	}

	//观察长时间没有请求之后重新有请求统计是否正确

	<-time.After(time.Millisecond * 100 * 30)
	fmt.Println("after 3s", win.Sum())
	win.Add(666)
	fmt.Println(win.Sum(), win)
}
