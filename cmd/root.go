/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : root
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lsh/initialize"
	"lsh/lib"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lsh [command] [args]",
	Short: "lsh is a CLI tool for managing file comments",
	Run: func(cmd *cobra.Command, args []string) {
		// 命令行解析
		if len(args) == 0 {
			lib.GetFileList(initialize.RuntimeInfo.RuntimePath)
		} else if len(args) == 1 {
			// 只有一个参数则为路径或者命令
			switch args[0] {
			case "help":
				help := `lsh is a CLI tool for managing file comments.
		Powered by Hyperia.

Usage:
	lsh			show file list in current path.
	lsh [path]		show file list in path.
	lsh [path] [command]	execute command in path.

Commands:
	add [comment]	add or update a comment to a file.
	del		delete a comment from a file.
	help		view usage help.
	version		view version.
	show 		show hidden files

example:
	lsh /home/user add "this is a comment"		add a comment to a file.
	lsh /home/user del				delete a comment from a file.`
				fmt.Println(help)
				return
			case "version":
				fmt.Println("lsh version 0.0.1 (2023-06-01) Beta")
				return
			case "show":
				initialize.RuntimeInfo.ShowHidden = true
				lib.GetFileList(initialize.RuntimeInfo.RuntimePath)
				return
			}
			// 校验路径是否正确
			if lib.CheckPathExist(args[0]) {
				lib.GetFileList(args[0])
			} else {
				fmt.Println("path or command is error !")
			}
		} else if len(args) >= 2 {
			// 检查args[0]是否为路径
			if lib.CheckPathExist(args[0]) {
				// 匹配命令
				switch args[1] {
				case "add":
					// 检查是否存在args[2]
					if len(args) != 3 {
						fmt.Printf("add must with comment !\n")
						return
					}
					err := lib.SetAttr(args[0], "user.comment", args[2])
					if err != nil {
						fmt.Printf("add comment %s to %s failed\n", args[2], args[0])
					} else {
						fmt.Printf("add comment %s to %s success\n", args[2], args[0])
					}
				case "del":
					err := lib.DelAttr(args[0], "user.comment")
					if err != nil {
						fmt.Printf("delete comment from %s failed\n", args[0])
					} else {
						fmt.Printf("delete comment from %s success\n", args[0])
					}
				default:
					fmt.Printf("command %s is error !\n", args[1])
				}
			} else {
				fmt.Println("path is not exist !")
			}

		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
