package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	homePath, _ := os.UserHomeDir()

	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(homePath, ".nano-claude-code"))

	BindEnv()

	err := viper.ReadInConfig()

	// TODO: 文件不存在，启动创建的引导

	return err
}
