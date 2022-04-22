package main

import (
	"fmt"
	"v2free-daily-attendance-go/config"
	signIn "v2free-daily-attendance-go/remote/signIn"
)

func main() {
	config.Init()
	msg := signIn.SignIn()
	fmt.Println("=======", msg, "========")

}
