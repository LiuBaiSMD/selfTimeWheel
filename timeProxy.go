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
	cancel    bool    //
}

func (tp *TimerProxy) Tick() {
	if tp.CheckValid(){
		tp.delayFunc()
	}
}

// 停止自己的定时任务，将任务的cancel标志位置为true
func (tp *TimerProxy) Cancel() {
	tp.cancel = true
}

// 判断定时任务是否有效
func (tp *TimerProxy) CheckValid() bool {
	return !tp.cancel
}
