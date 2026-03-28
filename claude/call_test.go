package claude

import (
	"os"
	"testing"
)

func newTestClient() *ClaudeClient {
	client, _ := NewClient(os.Getenv("baseurl"), os.Getenv("apikey"))
	return client
}

func TestCall(t *testing.T) {
	client := newTestClient()
	err := client.CallStream("claude-sonnet-4-6", []Message{
		{
			Role:    ClaudeMessageRoleUser,
			Content: SingleStringMessage("你好"),
		},
	}, func(m Message) bool {
		return true
	})
	if err != nil {
		t.Error(err.Error())
	}

}
