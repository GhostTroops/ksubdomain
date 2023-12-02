package go_utils

import (
	"log"
	"runtime"
	"strings"
	"sync"
)

type RegFuncs struct {
	FuncList []func()
}

// 并发执行所有func，并等待他们执行完
func WaitFunc(a ...func()) {
	var wg sync.WaitGroup
	WaitFunc4Wg(&wg, a...)
	wg.Wait()
}

// 并行执行方法，并将使用 wg 计数器
func WaitFunc4Wg(wg *sync.WaitGroup, a ...func()) {
	for _, x := range a {
		wg.Add(1)
		go func(cbk func()) {
			defer wg.Done()
			cbk()
		}(x)
	}
}

// 并行执行方法，并将使用 wg 计数器
// 同时传入参数parms
func WaitFunc4WgParms(wg *sync.WaitGroup, parms []any, a ...func(x ...any)) {
	for _, x := range a {
		wg.Add(1)
		go func(cbk func(...any)) {
			defer wg.Done()
			cbk(parms...)
		}(x)
	}
}

// 注册
func (r *RegFuncs) RegFunc(fn func()) {
	r.FuncList = append(r.FuncList, fn)
}

// tick 检测
var TickFunc = new(RegFuncs)
var ReleaseFunc = new(RegFuncs)

// 串行tick
func (r *RegFuncs) DoFunc() {
	for _, c := range r.FuncList {
		c()
	}
}

// Catch Panic
//
//	in your func: defer CatchPanic()
func CatchPanic() {
	if o := recover(); nil != o {
		log.Println(o)
	}
}

// 将该方法放到方法中运行，就可以打印出所有调用该方法的链路出来
func PrintCaller() {
	var i = 0
	for {
		i++
		if pc, file, line, ok := runtime.Caller(i); ok {
			fc := runtime.FuncForPC(pc)
			log.Printf("<-%s %s file:%s (line:%d)\n", strings.Repeat(":", i-1), fc.Name(), file, line) // , runtime.CallersFrames([]uintptr{pc})
			if "main.main" == fc.Name() {
				break
			}
		} else {
			break
		}
	}
}
