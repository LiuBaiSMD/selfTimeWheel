/*
@Author: liubai
@Date: 2021/5/21 下午5:38
@Desc: use for what
*/

package main

import (
	"fmt"
	"time"
)

type TimeWheel struct {
	timers       []*TimerProxy
	tickInterval time.Duration
	nowTime      int64
	exit         chan interface{}
}

//var timeWheel TimeWheel

func init() {
	// 启动定时
	fmt.Println("timeWheelinit")
	//timeWheel = TimeWheel{}
	//go timeWheel.Start()
	//defer timeWheel.Stop()
}

func (this *TimeWheel) Start() {
	//时间轮振动
	ticker := time.NewTicker(1 * this.tickInterval)
	defer ticker.Stop()
	for {
		this.nowTime = time.Now().UnixNano()
		select {
		case <-ticker.C:
			for {
				index, proxy := this.findProxy()
				if proxy == nil {
					break
				}
				fmt.Println("index:", index)
				this.delProxy(index)
				go proxy.Tick()
			}
		case <-this.exit:
			{
				fmt.Println("exit chan")
				return
			}
		}
	}

}

func (this *TimeWheel) Stop() {
	fmt.Println("TimeWheel over")
	close(this.exit)
}

func (this *TimeWheel) AddTimer(delay float64, delayFunc func()) *TimerProxy{
	nowTime := time.Now().UnixNano()
	fmt.Println("AddTimer: ", nowTime, int64(nanoTimerMultiTest*delay))
	timer := &TimerProxy{
		delay:     delay,
		delayFunc: delayFunc,
		tickTime:  nowTime + int64(nanoTimerMulti*delay), //time.Now().UnixNano()+int64(nanoTimerMultiTest*delay),
	}
	this.timers = append(timeWheel.timers, timer)
	return timer
}

func (this *TimeWheel) findProxy() (int, *TimerProxy) {
	for i, proxy := range this.timers {
		if proxy.tickTime <= this.nowTime {
			return i, proxy
		}
	}
	return -1, nil
}

func (this *TimeWheel) delProxy(index int) int {
	this.timers = append(this.timers[0:index], this.timers[index+1:]...)
	//fmt.Println("delAfter: ", this.timers)
	return index
}
