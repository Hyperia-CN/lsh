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
	"os"
	"path/filepath"
	"strings"
)

// 格式化输出
// 模拟ls命令对文件进行格式化输出、根据文件类型进行颜色区分
// 接收参数：文件结构体列表

// PrintFileList ls格式输出
func PrintFileList(fileList []*fileStruct) {
	// 遍历文件列表
	for _, file := range fileList {
		formatStdOut(file)
	}
	printFileNames(fileList)
}

// PrintTree 树形输出
func PrintTree(path string, prefix string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// 遍历文件列表
	for i, file := range files {
		var (
			isLast       = i == len(files)-1
			nestedPath   = filepath.Join(path, file.Name())
			nestedPrefix string
			f            = &fileStruct{}
		)

		// 初始化文件结构体
		f.init(file.Name(), nestedPath)

		// 判断是否为隐藏文件
		// 处理隐藏文件
		if f.fileType == "hiddenDir" || f.fileType == "hiddenFile" {
			// 判断是否显示隐藏文件
			if !initialize.RuntimeInfo.ShowHidden {
				continue // 跳过隐藏文件
			}
		}

		// 处理是否显示无注释文件
		if !configs.UserConfigs.ShowNoComment["tree"] {
			if f.comment == "" {
				continue
			}
		}

		fileList = append(fileList, f)

		// 格式化输出
		formatStdOut(f)

		// 添加树形结构
		if isLast {
			fmt.Printf("%s└── %s\n", prefix, f.color)
			nestedPrefix = prefix + "    "
		} else {
			fmt.Printf("%s├── %s\n", prefix, f.color)
			nestedPrefix = prefix + "│   "
		}

		// 递归调用，打印子目录
		if err := PrintTree(nestedPath, nestedPrefix); err != nil {
			return err
		}
	}

	return nil
}

func formatStdOut(file *fileStruct) {
	// 根据操作系统类型、文件类型对fileStruct中的name进行格式化并修改fileStruct中的name
	// 接收参数：file *fileStruct

	// 是否添加目录标识符
	if configs.UserConfigs.DirIndicator {
		// 修正BUG: 忘记判断是否为目录
		if file.fileType == "dir" || file.fileType == "hiddenDir" {
			file.name += configs.UserConfigs.DirIndicatorStr
		}
	}

	// 有注释的文件加上注释内容
	if file.comment != "" {
		file.name += configs.UserConfigs.CommentConnector + file.comment
	}

	// 根据文件类型对文件名进行颜色格式化
	// 检查 file.fileType 是否在 configs.UserConfigs.HighlightScheme 中
	if _, ok := configs.UserConfigs.HighlightScheme[file.fileType]; ok {
		file.color = echoColor(file.name, configs.UserConfigs.HighlightScheme[file.fileType])
	} else {
		file.color = echoColor(file.name, "white")
	}
}

func echoColor(str string, color string) string {
	// 对字符串进行颜色格式化
	// 判断是否为有效颜色
	if _, ok := configs.UserConfigs.ColorMap[color]; ok {
		//return fmt.Sprintf("\033[%sm%s\033[0m", initialize.RuntimeInfo.ColorMap[color], str)
		// 判断有无注释，有注释的话对文件名和注释分别进行颜色格式化
		if strings.Contains(str, configs.UserConfigs.CommentConnector) && configs.UserConfigs.CommentAlone {
			tmpStr := ""
			splitStr := strings.Split(str, configs.UserConfigs.CommentConnector)
			tmpStr += addColor(splitStr[0], color)
			// 连接注释连接符
			tmpStr += configs.UserConfigs.CommentConnector
			// 给注释进行颜色渲染
			tmpStr += addColor(splitStr[1], configs.UserConfigs.CommentColor)
			// 给文件名进行颜色渲染
			return tmpStr
			// 连接注释连接符
			// 给注释进行颜色渲染
		} else {
			return addColor(str, color)
		}
	} else {
		return str
	}
}

func addColor(str string, color string) string {
	tmpStr := initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["start"]
	tmpStr += str
	tmpStr += initialize.RuntimeInfo.ColorCode[initialize.RuntimeInfo.OS]["end"]
	return fmt.Sprintf(tmpStr, configs.UserConfigs.ColorMap[color])
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

		// 处理是否显示无注释文件
		if !configs.UserConfigs.ShowNoComment["ls"] {
			if file.comment == "" {
				continue
			}
		}

		// 处理注释文件输出样式
		if strings.Contains(file.name, configs.UserConfigs.CommentConnector) {
			switch configs.UserConfigs.CommentOutput {
			case "head":
				// 在头部输出有注释的文件
				fmt.Println(file.color)
				continue
			case "tail":
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

	// 当输出模式为 tail 时，输出有注释的文件
	if configs.UserConfigs.CommentOutput == "tail" {
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

// PrintError 打印错误信息
func PrintError(errMessage string, args ...[]string) {
	// 如果有args则先格式化错误信息
	if len(args) > 0 {
		for _, arg := range args {
			//errMessage += arg + " "
			for _, str := range arg {
				errMessage += str + " "
			}
		}
	}
	// 打印错误信息
	fmt.Println(addColor(errMessage, "red"))
}
