// package config 负责配置信息
package config

import (
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置viper实例冲突
)

// viper 库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 在动态生成配置信息
var ConfigFuncs map[string]interface{}

func init() {

	// 1. 初始化Viper库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//                "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	// 3. 环境变量配置文件查找的路径，相对于main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分Go的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)
	// 2. 注册配置信息
	loadConfig()
}
