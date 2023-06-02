/*
@Time : 2023/6/30 15:22
@Author : Hyperia
@File : parameters
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package cmd

type Params struct {
	path    string // 路径
	command string // 命令
	comment string // 注释
}

type Command struct {
	name  string // Command name
	short string // short Command
	run   func() // Command run function
}

// Commands 命令列表
type Commands []*Command

// Args 参数
var Args Params

// FindCommand 查找命令
func (c *Commands) FindCommand(args []string) *Command {
	// 首先检查是否有参数
	if len(args) == 0 {
		return nil
	}

	// 检查是存在命令参数
	for _, cmd := range *c {
		for _, arg := range args {
			if cmd.name == arg || cmd.short == arg {
				return cmd
			}
		}
	}

	// 没有找到匹配的命令
	return nil
}

// RunParams 执行参数
// 通过对Args的操作，可以修改参数
func (p *Params) RunParams(f func(p *Params) *Params) {
	Args = *f(p)
}

// Register 注册命令
func (c *Commands) Register(command *Command) {
	*c = append(*c, command)
}

// Execute 执行命令
func (c *Commands) Execute() {
	// 执行命令
	for _, command := range *c {
		if command.name == Args.command || command.short == Args.command {
			command.run()
		}
	}
}
