package cli

import "fmt"

var User = userCli{}

type userCli struct{}

// Test cli test
func (u userCli) Test() {
	fmt.Println("cli test...")
}
