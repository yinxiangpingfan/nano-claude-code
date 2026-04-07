package config

import "github.com/spf13/viper"

// 和环境变量绑定
func BindEnv() {
	viper.SetEnvPrefix("ncc")
	viper.BindEnv("llm.apikey")
	viper.BindEnv("llm.baseurl")
	viper.BindEnv("llm.model")
}
