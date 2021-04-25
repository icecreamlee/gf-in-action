package cron

import (
	"fmt"
)

var User = userCron{}

type userCron struct{}

// Test cron test
func (u userCron) Test() {
	fmt.Println("cron test...")
}
