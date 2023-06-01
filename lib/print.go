/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : print
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package lib

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"lsh/initialize"
	"math"
	"strings"
)

// 格式化输出
// 模拟ls命令对文件进行格式化输出、根据文件类型进行颜色区分
// 接收参数：文件结构体列表

func PrintFileList(fileList []*fileStruct) {
	// 遍历文件列表
	for _, file := range fileList {
		formatStdOut(file)
	}
	printFilenames(fileList)
}

func formatStdOut(file *fileStruct) {
	// 根据操作系统类型、文件类型对fileStruct中的name进行格式化并修改fileStruct中的name
	// 接收参数：file *fileStruct
	if initialize.RuntimeInfo.OS == "linux" || initialize.RuntimeInfo.OS == "darwin" {
		// 将文件夹特征添加到输出字符串中
		if file.fileType == "dir" || file.fileType == "hiddenDir" {
			file.name += "/"
		}
	} else if initialize.RuntimeInfo.OS == "windows" {
		// 将文件夹特征添加到输出字符串中
		if file.fileType == "dir" || file.fileType == "hiddenDir" {
			file.name += "\\"
		}
	}

	// 有注释的文件加上注释内容
	if file.comment != "" {
		file.name += "<--" + file.comment
	}

	// 根据文件类型对输出字符串进行颜色格式化
	if file.fileType == "dir" || file.fileType == "hiddenDir" {
		file.color = addColor(file.name, "cyan")
	} else if file.fileType == "hiddenFile" || file.fileType == "file" {
		file.color = addColor(file.name, "white")
	} else if file.fileType == "executable" {
		file.color = addColor(file.name, "red")
	} else {
		file.color = addColor(file.name, "white")
	}
}

func addColor(str string, color string) string {
	// 对字符串进行颜色格式化
	// 判断是否为有效颜色
	if _, ok := initialize.RuntimeInfo.ColorMap[color]; ok {
		//return fmt.Sprintf("\033[%sm%s\033[0m", initialize.RuntimeInfo.ColorMap[color], str)
		// 判断有无注释，有注释的话对文件名和注释分别进行颜色格式化
		if strings.Contains(str, "<--") {
			splitStr := strings.Split(str, "<--")
			return fmt.Sprintf("\033[%sm%s\033[0m<--\033[%sm%s\033[0m", initialize.RuntimeInfo.ColorMap[color], splitStr[0], initialize.RuntimeInfo.ColorMap["yellow"], splitStr[1])
		} else {
			return fmt.Sprintf("\033[%sm%s\033[0m", initialize.RuntimeInfo.ColorMap[color], str)
		}
	} else {
		return str
	}
}

func printFilenames(fileList []*fileStruct) {
	visibleFiles := make([]*fileStruct, 0) // 创建一个新的slice，用于存储非隐藏文件
	maxFilenameLength := 0
	for _, file := range fileList {
		if file.fileType == "hiddenDir" || file.fileType == "hiddenFile" {
			// 判断是否显示隐藏文件
			if !initialize.RuntimeInfo.ShowHidden {
				continue // 跳过隐藏文件
			}
		}
		visibleFiles = append(visibleFiles, file) // 向新的slice中添加非隐藏文件
		if length(file.name) > maxFilenameLength {
			maxFilenameLength = length(file.name)
		}
	}
	// 计算每行每列最多可以显示的文件数
	numColumns := int(math.Floor(float64(initialize.RuntimeInfo.TerminalWidth) / float64(maxFilenameLength+1)))
	numRows := int(math.Ceil(float64(len(visibleFiles)) / float64(numColumns)))
	// 按行打印文件名
	for i := 0; i < numRows; i++ {
		rowFiles := visibleFiles[i*numColumns : int(math.Min(float64((i+1)*numColumns), float64(len(visibleFiles))))]
		rowText := ""
		for _, file := range rowFiles {
			rowText += file.color + strings.Repeat(" ", maxFilenameLength-length(file.name)+1)
		}
		fmt.Println(rowText)
	}
}

// 返回字符串的终端显示宽度
func length(str string) int {
	width := 0
	for _, c := range str {
		// 使用go-runewidth包计算字符串的终端显示宽度
		width += runewidth.RuneWidth(c)
	}
	return width
}
