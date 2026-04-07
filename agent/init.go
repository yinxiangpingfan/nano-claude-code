package agent

import (
	"github.com/TIC-DLUT/nano-claude-code/claude"
	"github.com/spf13/viper"
)

type Agent struct {
	apiClient *claude.ClaudeClient
	tools     []claude.Tool
}

func NewAgent() (*Agent, error) {
	apiClient, err := claude.NewClient(viper.GetString("llm.baseurl"), viper.GetString("llm.apikey"))
	if err != nil {
		return nil, err
	}

	newAgent := &Agent{
		apiClient: apiClient,
	}

	newAgent.LoadTools()

	return newAgent, nil
}
