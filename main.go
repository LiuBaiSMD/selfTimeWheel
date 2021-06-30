/*
@Author: liubai
@Date: 2021/5/21 下午3:57
@Desc: use for what
*/

package main

import (
	"fmt"
	"time"
	timeWheelTools "timeWheelTest/timeWheel"
)

var tw *timeWheelTools.TimeWheel

func main() {
	// 初始化时间轮
	over := make(chan int, 1)
	//timeWheel = timeWheelTools.TimeWheel{
	//	[]*TimerProxy{},
	//	time.Millisecond * 100,
	//	time.Now().UnixNano(),
	//	make(chan interface{}),
	//}
	tw = timeWheelTools.NewTimeWheel(time.Second)
	go tw.Start()
	defer tw.Stop()
	l := []int{1, 2, 3, 4, 5}
	delayFunc := func() {
		fmt.Println("timer trigger print l:-------> ", l)
	}
	timerProxy := tw.AddTimer(0.5, delayFunc)
	fmt.Println("print :", l)
	l = l[:len(l)-1]
	//模拟定时器取消
	exitFunc := func() {
		over <- 1
	}
	tw.AddTimer(2, exitFunc)
	fmt.Println("timerProxy: ", timerProxy, exitFunc)
	select {
	case <-over:
		return
	}
}
