package main

import (
	"fmt"
	_ "v2free-daily-attendance-go/config"
	signIn "v2free-daily-attendance-go/remote/signIn"
)

func main() {
	msg := signIn.SignIn()
	fmt.Println("=======", msg, "========")
}
