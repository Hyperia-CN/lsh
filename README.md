# lsh
该工具可以用于对 ls 的一个功能性拓展。

基于对文件元数据的修改，实现了对文件进行注释后无论对文件进行移动、复制、编辑都能查看到对文件的注释。

## Asciicast
[![asciicast](https://asciinema.org/a/VqC8END8aIy66lIbE0JA9nKSc.svg)](https://asciinema.org/a/VqC8END8aIy66lIbE0JA9nKSc)

## 项目灵感

[//]: # (Github有许多优秀的开源的项目，但是因为使用不频繁，项目名称又不够直观，所以经常会忘记项目的名称，导致每次都要去搜索，浪费时间。之前我都是在很多工具的根目录建立一个 README.md 文件进行标注对应的项目是什么。所以导致我每次都要去阅读一遍 README.md 文件，浪费时间。所以我想到了能不能在文件上直接进行标注，这样就不用去阅读 README.md 文件了。5.30日突发奇想要不我写一个工具来实现这个功能，于是就有了这个项目。)

[//]: # ()
[//]: # (我比较熟悉Python和Java但是为什么选择Go语言？因为Go语言的跨平台性，我可以在Windows、Linux、Mac上都能使用这个工具。所以我看了几个小时的Go语言的文档，然后就有了这个项目。)

[//]: # ()
[//]: # (当然我对GO语言的开发实践还没有经验，如果各位有好的工程结构建议或者解决方案或者想法，欢迎各位提出issue，我会尽快解决。)

Github有许多优秀的开源的项目，但是因为使用不频繁，项目名称又不够直观，导致每次都要去搜索，浪费时间。为了解决这个问题，我想开发一个工具来直接在文件上标注对应的项目名称。这个工具将提供易于使用的标注方法，以便用户可以快速地找到他们所需的项目。

我选择使用Go语言来开发这个工具，因为它具有跨平台性，可以在多个操作系统上使用。虽然我对Go语言的开发实践还没有经验，但我花了几个小时阅读Go语言的文档，并且我相信我可以通过不断地学习和实践来提高自己的技能和经验。

如果你有好的工程结构建议、解决方案或者想法，欢迎提出issue。我非常欢迎任何形式的反馈和建议，并将尽快地处理它们。同时，如果你在使用这个工具时遇到任何问题或者有任何改进的建议，也欢迎随时联系我。

具体来说，我希望得到以下方面的帮助和建议：

- 工程结构建议：我希望得到一些关于如何更好地组织代码和文件结构的建议，以便使代码更易于维护和扩展。
- 性能优化建议：如果你有关于如何提高工具的性能的建议。
- 交互建议：如果你有关于如何改进工具的交互的建议。
- 功能性建议：如果你有关于如何改进工具的功能的建议。
- 语言本土化建议：如果你有关于如何改进工具的语言本土化的建议。
- 解决方案建议：如果你有关于如何解决某些问题或优化的建议。
- 任何其他建议和反馈：如果你有任何其他方面的建议和反馈，也欢迎随时联系我。

我将非常感谢你提供的帮助和建议，并将尽快考虑并采纳它们。最后，我希望这个工具能够帮助更多的人更好地管理和使用他们的项目，感谢你的支持！

## 功能介绍
1. 支持对指定的「文件/文件夹」「创建/查看」**注释**
2. 支持基本的ls文件查看功能
3. 高亮显示

## 安装
### Windows install
1. 下载[release](
2. 将下载的文件放入到环境变量中
3. 打开cmd，输入lsh version，如果出现版本号，则安装成功

### Linux install
1. 下载 [release]()
2. 将下载的文件放入到环境变量中
3. 打开终端，输入lsh version，如果出现版本号，则安装成功

### Mac install
1. 下载 [release]()
2. 将下载的文件放入到环境变量中
3. 打开终端，输入lsh version，如果出现版本号，则安装成功

## 使用帮助

```bash
    lsh // 查看当前文件夹下的文件包含注释
    lsh [path] // 查看指定文件夹下的文件包含注释
    lsh [path] add [comment] // 为指定文件夹添加注释
    lsh [path] del // 删除指定文件夹的注释
    lsh show // 显示隐藏文件
    lsh version // 查看版本号
    lsh help // 查看帮助
```

## TODO: 
- [ ] 代码重构、工程结构优化
- [ ] 适配 Windows、Linux
- [ ] 增加配置文件
- [ ] 增加自定义高亮配色
- [ ] 使用新的命令行参数解决方案
- [ ] 头脑风暴中...

## License
yd lsh is released under the MIT license. See [MIT](https://choosealicense.com/licenses/mit/)
