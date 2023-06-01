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
	// 终端颜色映射表
	ColorMap map[string]string
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
	// 初始化终端颜色映射表
	RuntimeInfo.ColorMap = map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	}
}

func terminalSize(w io.Writer) (int, int, error) {
	//GetFdInfo返回操作系统的文件描述符。文件并指示该文件是否表示终端。
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return 0, 0, fmt.Errorf("given writer is no terminal")
	}
	//GetWinsize根据指定的文件描述符返回窗口大小。
	winsize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(winsize.Width), int(winsize.Height), nil
}

// 旧的获取终端宽度方法
//func GetTerminalWidth() (int, error) {
//	var cmd *exec.Cmd
//	switch initialize.RuntimeInfo.OS {
//	case "linux":
//		cmd = exec.Command("stty", "size")
//	case "darwin":
//		cmd = exec.Command("stty", "size")
//	case "windows":
//		cmd = exec.Command("powershell", "Get-Host", "|", "Select-Object", "Width")
//	}
//	cmd.Stdin = os.Stdin
//	out, err := cmd.Output()
//	if err != nil {
//		return 0, err
//	}
//	width, err := strconv.Atoi(strings.Replace(strings.Split(string(out), " ")[1], "\n", "", -1))
//	if err != nil {
//		return 0, err
//	}
//	return width, nil
//}
