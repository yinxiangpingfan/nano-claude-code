package main

import "fmt"

func DirectRun() {
	if Message == "" {
		panic("message不能为空")
	}

	fmt.Println("开始处理", Message)

	MainAgent.ChatStream(Message, func(s string) {
		fmt.Print(s)
	})
}
