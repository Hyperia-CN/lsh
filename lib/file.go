/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : file
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package lib

import (
	"fmt"
	"lsh/initialize"
	"os"
	"strings"
)

// fileStruct 文件数据结构
type fileStruct struct {
	// 名称
	name string
	// color string
	color string
	// 路径
	path string
	// 注释
	comment string
	// 文件类型
	fileType string
}

// fileList 文件列表
var fileList []*fileStruct

// init 文件初始化
func (f *fileStruct) init(name string, path string) {
	f.name = name
	f.path = path
	f.fileType = f.checkFileType()
	f.comment, _ = GetAttr(f.path, "user.comment")
}

func (f *fileStruct) getData() {
	// 读取文件注释
}

// 文件类型判断
func (f *fileStruct) checkFileType() string {
	file, err := os.Stat(f.path)
	if err != nil {
		// handle error
		return ""
	}

	if file.IsDir() {
		// 判断是否为隐藏文件夹
		if file.Name()[0] == '.' {
			return "hiddenDir"
		} else {
			return "dir"
		}
	}

	fileName := file.Name()
	if fileName[0] == '.' {
		return "hiddenFile"
	}

	if file.Mode().IsRegular() {
		if initialize.RuntimeInfo.OS == "darwin" || initialize.RuntimeInfo.OS == "linux" {
			if file.Mode()&0111 != 0 {
				return "executable"
			}
		} else if initialize.RuntimeInfo.OS == "windows" {
			if strings.HasSuffix(strings.ToLower(fileName), ".exe") {
				return "executable"
			}
		}
	}
	return "file"
}

// GetFileList 获取指定路径下的文件列表
func GetFileList(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		f := fileStruct{}
		f.init(file.Name(), path+"/"+file.Name())
		fileList = append(fileList, &f)
	}
	PrintFileList(fileList)
}

// CheckPathExist 检查文件路径是否存在
func CheckPathExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
