package agent

import (
	"github.com/TIC-DLUT/nano-claude-code/claude"
	"github.com/spf13/viper"
)

func (a *Agent) ChatStream(message string, callback func(string)) {
	lastToolCallID := ""

	a.apiClient.CallStreamTools(viper.GetString("llm.model"), GetNowSystemPrompt(), []claude.Message{
		{
			Role:    claude.ClaudeMessageRoleUser,
			Content: claude.SingleStringMessage(message),
		},
	}, a.tools, func(m claude.Message) bool {
		switch m.Content.(type) {
		case claude.TextBlock:
			callback(m.Content.(claude.TextBlock).Text)
		case claude.ToolUseBlock:
			tooluse := m.Content.(claude.ToolUseBlock)
			if tooluse.ID != lastToolCallID {
				lastToolCallID = tooluse.ID
				callback("\n[tool_use] " + tooluse.Name + "\n")
			}
		}
		return true
	})
}
