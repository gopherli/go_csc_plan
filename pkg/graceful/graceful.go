package graceful

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var goRoutineCounter sync.WaitGroup

var stopChan chan struct{}

func init() {
	stopChan = make(chan struct{})
}

func Go(f func()) {
	goRoutineCounter.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// 获取调用栈信息
				buf := make([]byte, 4096)
				runtime.Stack(buf, false)
				content := "graceful Go recover panic:" + fmt.Sprintf("%s\n", r) + string(buf)
				fmt.Println("[ERR PANIC]", content)
			}
		}()
		defer goRoutineCounter.Done()
		f()
	}()
}

func Wait() {
	goRoutineCounter.Wait()
}

func Add(delta int) {
	goRoutineCounter.Add(delta)
}

func Done() {
	goRoutineCounter.Done()
}

func Stop() {
	close(stopChan)
}

func TimeInterval(t time.Duration, callback func()) {
	ticker := time.NewTicker(t)
	go func() {
		for {
			select {
			case <-ticker.C:
				Go(callback)
			case <-stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}
