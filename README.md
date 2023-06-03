# lsh
该工具可以用于对 ls 的一个功能性拓展。

基于对文件元数据的修改，实现了对文件进行注释后无论对文件进行移动、复制、编辑都能查看到对文件的注释。

## Asciicast
[![asciicast](https://asciinema.org/a/R5hpRnTbUeAYIQ7l4G2CVmS80.svg)](https://asciinema.org/a/R5hpRnTbUeAYIQ7l4G2CVmS80)

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
1. 支持对指定的「文件/文件夹」「创建/查看/删除」**注释**
2. 支持基本的ls文件查看功能
3. 高亮显示
4. 支持树形结构显示

## 安装
### Windows install -- 暂不支持
1. 下载 [release]()
2. 将下载的文件放入到环境变量中
3. 打开cmd，输入lsh version，如果出现版本号，则安装成功

### Linux install   -- 未测试
1. 下载 [release]()
2. 将下载的文件放入到环境变量中
3. 打开终端，输入lsh version，如果出现版本号，则安装成功

### Mac install
1. 下载 [release](https://github.com/Hyperia-CN/lsh/releases/download/1.0/lsh_darwin_x86_64)
2. 将下载的文件放入到环境变量中
3. 打开终端，输入lsh version，如果出现版本号，则安装成功

## 使用帮助

### 命令
```bash
    lsh // 查看当前文件夹下的文件包含注释
    lsh [path] // 查看指定文件夹下的文件包含注释
    lsh [path] add [comment] // 为指定文件夹添加注释
    lsh [path] del // 删除指定文件夹的注释
    lsh [path] show // 显示指定文件夹的隐藏文件
    lsh show // 显示隐藏文件
    lsh version // 查看版本号
    lsh help // 查看帮助
```

### 配置文件
```yaml
Windows配置文件路径：C:\Users<username>\.lshrc
Linux配置文件路径：/home/<username>/.lshrc  
Mac配置文件路径：/Users/<username>/.lshrc
```

- 自定义目录标识符
- 自定义注释连接符
- 自定义注释颜色
- 自定义高亮配色方案
- 自定义注释输出方式
- 自定义颜色映射

```yaml
shownocomment:  # 是否显示没有注释的文件
  ls: true   # 注意：默认调用ls命令，如果不想显示没有注释的文件，可以设置为false
  tree: false # 当使用tree参数时，是否显示没有注释的文件。建议设置为false否则会显示所有子文件夹的文件
dirindicator: true  # 是否添加目录标识符 「true/false」
dirindicatorstr: /  # 目录标识符 「string」
commentconnector: <-- # 注释连接符 「string」
commentalone: true  # 注释单独渲染颜色 「true/false」
commentcolor: yellow  # 注释颜色 「black/red/green/yellow/blue/magenta/cyan/white」
highlightscheme:  # 高亮配色方案 「default/monokai」
    dir: cyan # 目录
    executable: red # 可执行文件
    file: white # 文件
    hiddenDir: cyan # 隐藏目录
    hiddenFile: white # 隐藏文件
    other: white  # 其他
commentoutput: default  # 注释输出方式 「default/head/tail」
colormap: # 颜色映射 注意：颜色映射 id 为终端渲染的 ANSI 颜色 id
    black: "30"
    blue: "34"
    cyan: "36"
    green: "32"
    magenta: "35"
    red: "31"
    white: "37"
    yellow: "33"
```

## TODO: 
- [x] 增加配置文件
- [x] 增加自定义高亮配色
- [ ] 适配 Windows、Linux
- [x] 代码重构、工程结构优化
- [x] 使用新的命令行参数解决方案
- [ ] 增加 install 命令进行安装
- [ ] 增加 upgrade 命令进行升级
- [ ] 头脑风暴中...


## 更新日志
### v1.0.3 (2023-06-03) Beta
- 增加tree命令，可以查看指定目录下的文件树
- 增加配置文件自定义是否显示没有注释的文件

### v1.0.2 (2023-06-02) Beta
- 增加短参数支持
- show 命令增加对指定路径的支持
- 增加 head 和 end 命令可以自定义注释输出位置
- 放弃使用 cobra 解析命令行参数，自行抽象参数功能

### v1.0.0 (2023-06-01) Beta
- 增加自定义配置文件

### v0.0.1 (2023-05-31) Beta
- 实现基本功能

## License
lsh is released under the MIT license. See [MIT](https://choosealicense.com/licenses/mit/)

