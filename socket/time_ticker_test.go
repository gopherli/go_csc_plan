package socket

import (
	"fmt"
	"testing"
	"time"
)

func Test_TimeTicker(t *testing.T) {
	ClearDisconnectUserTimerStart()
}

func ClearDisconnectUserTimerStart() {
	TimeInterval(5*time.Second, func() { ClearDisconnectUser() })
}

func ClearDisconnectUser() {
	fmt.Println("clear")
}

func TimeInterval(t time.Duration, callback func()) *time.Ticker {
	ticker := time.NewTicker(t)
	go func() {
		for range ticker.C {
			go callback()
		}
	}()
	for {
		time.Sleep(time.Second * 1)
	}
	return ticker
}
