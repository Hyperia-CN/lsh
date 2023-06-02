/*
@Time : 2023/6/1 16:07
@Author : Hyperia
@File : configs
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"lsh/initialize"
	"os"
)

// Windows配置文件路径：用户配置文件：C:\Users<username>\.lshrc
// Linux配置文件路径：用户配置文件：/home/<username>/.lshrc
// Mac配置文件路径：用户配置文件：/Users/<username>/.lshrc

// UserConfigs 配置文件结构体
type configs struct {
	// 是否添加目录标识符
	DirIndicator bool // true | false
	// 目录标识符
	DirIndicatorStr string // "/" | "\\"
	// 注释连接符
	CommentConnector string // "--" | "++" | "##"
	// 注释是否单独渲染
	CommentAlone bool // true | false
	// 注释渲染颜色
	CommentColor string // "red" | "green" | "blue" | "yellow" | "magenta" | "cyan" | "white"
	// 高亮配色方案
	HighlightScheme map[string]string // {文件类型：颜色}
	// 注释打印样式
	CommentOutput string // "default" | "start" | "end"
	// 终端颜色映射表
	ColorMap map[string]string // {颜色：颜色值}
}

// 默认配置项，用于初始化配置文件
var defaultConfigs = configs{
	// 是否添加目录标识符
	DirIndicator: true,
	// 目录标识符
	DirIndicatorStr: "/",
	// 注释连接符
	CommentConnector: "<--",
	// 注释是否单独渲染
	CommentAlone: true,
	// 注释渲染颜色
	CommentColor: "yellow",
	// 高亮配色方案
	HighlightScheme: map[string]string{
		"file":       "white",
		"hiddenFile": "white",
		"dir":        "cyan",
		"hiddenDir":  "cyan",
		"executable": "red",
		"other":      "white",
	},
	CommentOutput: "default",
	ColorMap: map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	},
}

// UserConfigs 配置文件
var UserConfigs configs

// InitConfigs 初始化配置文件
func InitConfigs() {
	//UserConfigs = defaultConfigs
	// 检查配置文件是否存在
	if _, err := os.Stat(initialize.RuntimeInfo.ConfigPath); os.IsNotExist(err) {
		// 初始化配置文件
		UserConfigs = defaultConfigs
		saveConfigs()
	} else {
		// 存在则加载
		loadConfigs()
	}
}

// 加载配置文件
func loadConfigs() {
	// 加载配置文件
	data, err := os.ReadFile(initialize.RuntimeInfo.ConfigPath)
	if err != nil {
		fmt.Println("lsh：config file load failed.")
		panic(err)
	}
	err = yaml.Unmarshal(data, &UserConfigs)
	if err != nil {
		fmt.Println("lsh：config file load failed.")
		panic(err)
	}
}

// 保存配置文件
func saveConfigs() {
	// 保存配置文件
	data, err := yaml.Marshal(UserConfigs)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(initialize.RuntimeInfo.ConfigPath, data, 0644)
	if err != nil {
		panic(err)
	}
}
