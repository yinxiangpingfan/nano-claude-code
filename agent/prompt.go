package agent

import (
	"os"
	"strings"
	"time"
)

const (
	systemPrompt = `你是claude code，你需要调用工具帮助人们完成工作

当前时间是：{system_time}
当前工作地址是：{work_path}`
)

func GetNowSystemPrompt() string {
	systemTime := time.Now().Format("2006-01-02 15:04:05")
	workPath, err := os.Getwd()
	if err != nil {
		workPath = "unknown"
	}

	nowSystemPrompt := systemPrompt
	nowSystemPrompt = strings.ReplaceAll(nowSystemPrompt, "{system_time}", systemTime)
	nowSystemPrompt = strings.ReplaceAll(nowSystemPrompt, "{work_path}", workPath)

	return nowSystemPrompt
}
