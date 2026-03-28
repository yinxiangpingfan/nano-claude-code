package claude

type Message struct {
	Role    string `json:"role"`
	Content any    `json:"content"`
}

const (
	ClaudeMessageRoleUser      = "user"
	ClaudeMessageRoleAssistant = "assistant"
)

type SingleStringMessage string

type TextBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
