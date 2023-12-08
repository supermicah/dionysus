package wgroup

import (
	"sync"
	"time"
)

type WaitGroup struct {
	wg sync.WaitGroup
}

func (w *WaitGroup) Add(delta int) {
	w.wg.Add(delta)
}

func (w *WaitGroup) Done() {
	w.wg.Done()
}

// WaitSuccess 在规定时间内（`timeout`）是否等待成功
// 参数：
// timeout: 超时时间
func (w *WaitGroup) WaitSuccess(timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		w.wg.Wait()
	}()
	select {
	case <-c:
		return true // completed normally
	case <-time.After(timeout):
		return false // timed out
	}
}
