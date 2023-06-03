/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : env
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package initialize

import (
	"fmt"
	"github.com/moby/term"
	"io"
	"os"
	"runtime"
)

type runtimeInfo struct {
	// 系统类型
	OS string
	// 运行路径
	RuntimePath string
	// 显示隐藏文件 default: false
	ShowHidden bool
	// 终端宽度
	TerminalWidth int
	// 配置文件路径
	ConfigPath string
	// 终端添加颜色代码
	ColorCode map[string]map[string]string
	// 版本号
	Version string
	// 帮助信息
	Help string
}

var RuntimeInfo runtimeInfo

func Init() {
	// 初始化操作系统类型
	RuntimeInfo.OS = runtime.GOOS
	// 初始化运行目录
	pwd, err := os.Getwd()
	if err == nil {
		RuntimeInfo.RuntimePath = pwd
	} else {
		fmt.Println(err)
	}
	// 初始化是否显示隐藏文件
	RuntimeInfo.ShowHidden = false
	// 初始化终端宽度
	RuntimeInfo.TerminalWidth, _, _ = terminalSize(os.Stdout)
	// 初始化配置文件路径
	if RuntimeInfo.OS == "linux" || RuntimeInfo.OS == "darwin" {
		RuntimeInfo.ConfigPath = os.Getenv("HOME") + "/.lshrc"
	} else if RuntimeInfo.OS == "windows" {
		RuntimeInfo.ConfigPath = os.Getenv("USERPROFILE") + "/.lshrc"
	}

	// 初始化 Unix 系统终端添加颜色代码
	RuntimeInfo.ColorCode = map[string]map[string]string{
		"darwin": {
			"start": "\033[%sm",
			"end":   "\033[0m",
		},
		"linux": {
			"start": "\033[%sm",
			"end":   "\033[0m",
		},
		"windows": {
			"start": "",
			"end":   "",
		},
	}

	// 初始化版本号
	RuntimeInfo.Version = "v1.0.3 (2023-06-03) Beta"
	// 初始化帮助信息
	RuntimeInfo.Help = `lsh is a CLI tool for managing file comments.
		Powered by Hyperia.

Usage:
	lsh			show file list in current path.
	lsh [path]		show file list in path.
	lsh [path] [command]	execute Command in path.

Commands:
    command|short    args	description
	add|a	  [comment]	add or update a comment to a file.
	del|d			delete a comment from a file.
	show|s 			show hidden files.
	tree|T			show file tree.
	head|h			show comment in head.
	tail|t			show comment in tail.
	help|H			view usage help.
	version|v		view version.

Example:
	lsh /home/user add "comment"	add a comment to a file.
	lsh /home/user del		delete a comment from a file.
	lsh /home/user show		show hidden files.
	lsh /home/user tree		show file tree.`
}

func terminalSize(w io.Writer) (int, int, error) {
	// GetFdInfo 返回操作系统的文件描述符。文件并指示该文件是否表示终端。
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return 0, 0, fmt.Errorf("given writer is no terminal")
	}
	// GetWinsize 根据指定的文件描述符返回窗口大小。
	winsize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(winsize.Width), int(winsize.Height), nil
}
