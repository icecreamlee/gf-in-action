package cron

import (
	"fmt"
	"time"
)

var User = userCron{}

type userCron struct{}

// Test cron test
func (u userCron) Test() {
	defer fmt.Println("exit")
	fmt.Println("cron test...")
	time.Sleep(time.Hour)
}
