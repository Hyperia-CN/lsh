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
	"lsh/configs"
	"lsh/initialize"
	"lsh/lib"
	"os"
)

var lsCmd = Command{
	name: "ls",
	run: func() {
		lib.GetFileList(Args.path)
	},
}

var addCmd = Command{
	name:  "add",
	short: "a",
	run: func() {
		if Args.comment != "" {
			err := lib.SetAttr(Args.path, "user.comment", Args.comment)
			if err != nil {
				lib.PrintError(err.Error())
			} else {
				fmt.Printf("add comment \"%s\" to %s success\n", Args.comment, Args.path)
			}
		} else {
			lib.PrintError("comment is empty")
		}
	},
}

var delCmd = Command{
	name:  "del",
	short: "d",
	run: func() {
		err := lib.DelAttr(Args.path, "user.comment")
		if err != nil {
			lib.PrintError(err.Error())
		} else {
			fmt.Printf("del comment from %s success\n", Args.path)
		}
	},
}

var showCmd = Command{
	name:  "show",
	short: "s",
	run: func() {
		initialize.RuntimeInfo.ShowHidden = true
		lib.GetFileList(Args.path)
	},
}

var treeCmd = Command{
	name:  "tree",
	short: "T",
	run: func() {
		lib.GetFileTree(Args.path)
	},
}

var headCmd = Command{
	name:  "head",
	short: "h",
	run: func() {
		configs.UserConfigs.CommentOutput = "head"
		lib.GetFileList(Args.path)
	},
}

var endCmd = Command{
	name:  "tail",
	short: "t",
	run: func() {
		configs.UserConfigs.CommentOutput = "tail"
		lib.GetFileList(Args.path)
	},
}

var versionCmd = Command{
	name:  "version",
	short: "v",
	run: func() {
		fmt.Printf("lsh version %s\n", initialize.RuntimeInfo.Version)
	},
}

var helpCmd = Command{
	name:  "help",
	short: "H",
	run: func() {
		fmt.Println(initialize.RuntimeInfo.Help)
	},
}

func Init() {
	var commands Commands
	// 注册命令
	commands.Register(&versionCmd)
	commands.Register(&helpCmd)
	commands.Register(&lsCmd)
	commands.Register(&addCmd)
	commands.Register(&delCmd)
	commands.Register(&showCmd)
	commands.Register(&headCmd)
	commands.Register(&endCmd)
	commands.Register(&treeCmd)

	// 解析参数
	Args.RunParams(func(p *Params) *Params {
		// 去除执行文件名
		os.Args = os.Args[1:]
		// 首先检查是否有参数
		if len(os.Args) == 0 {
			// 没有参数则path为当前路径
			Args.path = initialize.RuntimeInfo.RuntimePath
			Args.command = "ls"
			return p
		} else {
			// 首先检查第一个参数是否为路径
			if lib.CheckPathExist(os.Args[0]) {
				Args.path = os.Args[0]
				if len(os.Args) >= 2 {
					// 如果有参数则检查是否为命令
					if cmd := commands.FindCommand(os.Args[1:2]); cmd != nil {
						Args.command = cmd.name
						// 增加验证，如果是add命令则检查是否有comment参数
						if len(os.Args) >= 3 && cmd.name == "add" {
							for i := 2; i < len(os.Args); i++ {
								Args.comment = Args.comment + os.Args[i] + " "
							}
							// 如果有comment参数但命令不是add则报错
						} else if len(os.Args) >= 3 {
							lib.PrintError("invalid comment：", os.Args)
							os.Exit(-1)
						} else {
							// 执行命令
							return p
						}
						return p
					} else {
						return p
					}
				} else {
					Args.command = "ls"
					return p
				}
			} else {
				// 如果第一个参数不是路径则设置路径为当前路径
				Args.path = initialize.RuntimeInfo.RuntimePath
				if len(os.Args) == 1 {
					// 如果不是路径则检查是否为命令
					if cmd := commands.FindCommand(os.Args); cmd != nil {
						Args.command = cmd.name
						return p
					} else {
						// 如果挤既不是路径也不是命令则报错
						lib.PrintError("invalid command or path：", os.Args)
						os.Exit(-1)
					}
				} else {
					// 如果第一个参数既不是路径也不是命令则报错
					lib.PrintError("invalid command or path：", os.Args)
					os.Exit(-1)
				}

			}
		}
		return p
	})
	// 执行命令
	commands.Execute()
}
