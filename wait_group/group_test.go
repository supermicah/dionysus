package wait_group

import (
	"testing"
	"time"
)

func TestWaitGroup1(t *testing.T) {
	wg := &WaitGroup{}
	for i := 0; i < 10; i++ {
		second := time.Duration(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * second)
		}()
	}
	success := wg.WaitSuccess(time.Second * 5)
	if success {
		t.Errorf("TestWaitGroup want: %t, got: %t", false, success)
	}
}
func TestWaitGroup2(t *testing.T) {
	wg := &WaitGroup{}
	for i := 0; i < 3; i++ {
		second := time.Duration(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * second)
		}()
	}
	success := wg.WaitSuccess(time.Second * 5)
	if !success {
		t.Errorf("TestWaitGroup want: %t, got: %t", true, success)
	}
}
