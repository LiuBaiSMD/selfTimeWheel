/*
@Author: liubai
@Date: 2021/5/21 下午4:02
@Desc: use for what
*/

package main

const (
	nanoTimerMulti     = 1000000000.0
	nanoTimerMultiTest = 1000.0
)

type TimerProxy struct {
	delay     float64 //延时长度
	delayFunc func()  //延时调用的方法
	tickTime  int64   //调用时间
	repeat    bool    //调用完毕后重新写入
}

func (tp *TimerProxy) Tick() {
	tp.delayFunc()
}

// 这个方法需要停止自己的定时任务
func (tp *TimerProxy) Cancle() {
	tp.delayFunc()
}
