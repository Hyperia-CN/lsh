/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : lsh
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package main

import (
	"lsh/cmd"
	"lsh/initialize"
)

func main() {
	initialize.Init()
	cmd.Execute()
}
