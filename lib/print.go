/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : print
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package lib

import "C"
import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"lsh/configs"
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
	printFileNames(fileList)
}

func formatStdOut(file *fileStruct) {
	// 根据操作系统类型、文件类型对fileStruct中的name进行格式化并修改fileStruct中的name
	// 接收参数：file *fileStruct
	if configs.UserConfigs.DirIndicator {
		file.name += configs.UserConfigs.DirIndicatorStr
	}

	// 有注释的文件加上注释内容
	if file.comment != "" {
		file.name += configs.UserConfigs.CommentConnector + file.comment
	}

	// 根据文件类型对文件名进行颜色格式化
	// 检查 file.fileType 是否在 configs.UserConfigs.HighlightScheme 中
	if _, ok := configs.UserConfigs.HighlightScheme[file.fileType]; ok {
		file.color = addColor(file.name, configs.UserConfigs.HighlightScheme[file.fileType])
	} else {
		file.color = addColor(file.name, "white")
	}
}

func addColor(str string, color string) string {
	// 对字符串进行颜色格式化
	// 判断是否为有效颜色
	if _, ok := configs.UserConfigs.ColorMap[color]; ok {
		//return fmt.Sprintf("\033[%sm%s\033[0m", initialize.RuntimeInfo.ColorMap[color], str)
		// 判断有无注释，有注释的话对文件名和注释分别进行颜色格式化
		if strings.Contains(str, configs.UserConfigs.CommentConnector) && configs.UserConfigs.CommentAlone {
			splitStr := strings.Split(str, configs.UserConfigs.CommentConnector)
			// 给文件名进行颜色渲染
			tmpStr := initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["start"]
			tmpStr += splitStr[0]
			tmpStr += initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["end"]
			// 连接注释连接符
			tmpStr += configs.UserConfigs.CommentConnector
			// 给注释进行颜色渲染
			tmpStr += initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["start"]
			tmpStr += splitStr[1]
			tmpStr += initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["end"]
			// 给文件名进行颜色渲染
			tmpStr = fmt.Sprintf(
				tmpStr,
				configs.UserConfigs.ColorMap[color],
				configs.UserConfigs.ColorMap[configs.UserConfigs.CommentColor])
			return tmpStr
			// 连接注释连接符
			// 给注释进行颜色渲染
		} else {
			tmpStr := initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["start"]
			tmpStr += str
			tmpStr += initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["end"]
			return fmt.Sprintf(tmpStr, configs.UserConfigs.ColorMap[color])
		}
	} else {
		return str
	}
}

func printFileNames(fileList []*fileStruct) {
	visibleFiles := make([]*fileStruct, 0) // 创建一个新的slice，用于存储非隐藏文件
	maxFilenameLength := 0
	for _, file := range fileList {
		// 处理隐藏文件
		if file.fileType == "hiddenDir" || file.fileType == "hiddenFile" {
			// 判断是否显示隐藏文件
			if !initialize.RuntimeInfo.ShowHidden {
				continue // 跳过隐藏文件
			}
		}

		// 处理注释文件输出样式
		if strings.Contains(file.name, configs.UserConfigs.CommentConnector) {
			switch configs.UserConfigs.CommentOutput {
			case "start":
				// 在头部输出有注释的文件
				fmt.Println(file.color)
				continue
			case "end":
				// 在尾部输出有注释的文件
				continue
			case "default":
				// 正常按照顺序输出
			}
		}

		// 处理文件名长度
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

	// 当输出模式为end时，输出有注释的文件
	if configs.UserConfigs.CommentOutput == "end" {
		for _, file := range fileList {
			if strings.Contains(file.name, configs.UserConfigs.CommentConnector) {
				fmt.Println(file.color)
			}
		}
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
