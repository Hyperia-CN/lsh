/*
@Time : 2023/6/1 16:07
@Author : Hyperia
@File : configs
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"lsh/initialize"
	"os"
)

// Windows配置文件路径：用户配置文件：C:\Users<username>\.lshrc
// Linux配置文件路径：用户配置文件：/home/<username>/.lshrc
// Mac配置文件路径：用户配置文件：/Users/<username>/.lshrc

// UserConfigs 配置文件结构体
type configs struct {
	// 默认是否显示无注释文件
	ShowNoComment map[string]bool // {tree: false, ls: true}
	// 是否添加目录标识符
	DirIndicator bool // true | false
	// 目录标识符
	DirIndicatorStr string // "/" | "\\"
	// 注释连接符
	CommentConnector string // "--" | "++" | "##"
	// 注释是否单独渲染
	CommentAlone bool // true | false
	// 注释渲染颜色
	CommentColor string // "red" | "green" | "blue" | "yellow" | "magenta" | "cyan" | "white"
	// 后缀名分类
	SuffixClassify map[string][]string // {文件类型：[后缀名]}
	// 高亮配色方案
	HighlightScheme map[string]string // {文件类型：颜色}
	// 注释打印样式
	CommentOutput string // "default" | "head" | "tail"
	// 终端颜色映射表
	ColorMap map[string]string // {颜色：颜色值}
}

// 默认配置项，用于初始化配置文件
var defaultConfigs = configs{
	// 默认是否显示无注释文件
	ShowNoComment: map[string]bool{
		"tree": false,
		"ls":   true,
	},
	// 是否添加目录标识符
	DirIndicator: true,
	// 目录标识符
	DirIndicatorStr: "/",
	// 注释连接符
	CommentConnector: "<--",
	// 注释是否单独渲染
	CommentAlone: true,
	// 注释渲染颜色
	CommentColor: "yellow",
	// 高亮配色方案
	HighlightScheme: map[string]string{
		"file":       "white",
		"hiddenFile": "white",
		"dir":        "cyan",
		"hiddenDir":  "cyan",
		"executable": "red",
		"code":       "green",
		"other":      "white",
	},
	CommentOutput: "default",
	ColorMap: map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	},
	// 后缀名分类
	SuffixClassify: map[string][]string{
		"image":   {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".ico", ".webp", ".tif", ".tiff", ".pcx", ".tga", ".exif", ".fpx", ".svg", ".psd", ".cdr", ".pcd", ".dxf", ".ufo", ".eps", ".ai", ".raw", ".WMF"},
		"video":   {".avi", ".mov", ".rmvb", ".rm", ".flv", ".mp4", ".3gp", ".mpeg", ".mpg", ".dat", ".mkv", ".wmv", ".asf", ".asx", ".vob"},
		"audio":   {".mp3", ".aac", ".wav", ".wma", ".cda", ".flac", ".m4a", ".mid", ".mka", ".mp2", ".mpa", ".mpc", ".ape", ".ofr", ".ogg", ".ra", ".wv", ".tta", ".ac3", ".dts"},
		"doc":     {".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".wps", ".rtf", ".csv", ".pps", ".et", ".dps", ".pot", ".pps", ".xlt", ".xlw", ".dot", ".xml", ".html", ".htm", ".mht", ".mhtml", ".chm", ".epub", ".mobi", ".fb2", ".djvu", ".ps", ".xps", ".oxps", ".dotx", ".dotm", ".docm", ".docm", ".dotm", ".xlsb", ".xlsm", ".xltx", ".xltm", ".xlam", ".pptm", ".potm", ".ppam", ".ppsm", ".sldm", ".thmx", ".xps", ".oxps"},
		"archive": {".zip", ".rar", ".7z", ".arj", ".z", ".gz", ".bz2", ".tar", ".iso", ".cab", ".jar", ".ace", ".lzh", ".uue", ".gzip", ".7zip"},
		"code":    {".go", ".py", ".java", ".c", ".cpp", ".h", ".hpp", ".cs", ".php", ".js", ".ts", ".css", ".scss", ".less", ".html", ".htm", ".json", ".xml", ".yaml", ".yml", ".ini", ".toml", ".conf", ".sh", ".bat", ".cmd", ".ps1", ".vbs", ".lua", ".rb", ".pl", ".swift", ".scala", ".coffee", ".dart", ".erl", ".hs", ".lisp", ".clj", ".groovy", ".as", ".m", ".mm", ".php", ".py"},
		"font":    {".ttf", ".ttc", ".otf", ".woff", ".woff2", ".eot", ".fon"},
		"exe":     {".exe", ".msi", ".bat", ".cmd", ".com", ".reg", ".vb", ".vbs", ".vbe", ".js", ".jse", ".ws", ".wsf", ".wsc", ".wsh", ".ps1", ".ps1xml", ".ps2", ".ps2xml", ".psc1", ".psc2", ".msh", ".msh1", ".msh2", ".mshxml", ".msh1xml", ".msh2xml", ".scf", ".lnk", ".inf", ".reg", ".dll", ".sys", ".drv", ".cpl", ".ocx", ".ax", ".acm", ".ax", ".cpl", ".drv", ".efi", ".exe", ".mui", ".ocx", ".scr", ".sys", ".tsp", ".dll", ".exe", ".mui", ".ocx", ".scr", ".sys", ".tsp", ".efi", ".msc", ".msp", ".msu", ".diagcab", ".diagpkg", ".appx", ".appxbundle", ".appxupload", ".msix", ".msixbundle", ".cer", ".cert", ".crt", ".der", ".p7b", ".p7r", ".spc", ".sst", ".stl", ".pfx", ".p12", ".pem", ".key", ".asc", ".pgp", ".gpg", ".odt", ".ods", ".odp", ".odg", ".odc", ".odb", ".odf", ".wpd", ".xls", ".xlsx", ".ppt", ".pptx", ".vsd", ".vsdx", ".vsdm", ".vsd", ".vss", ".vst", ".vdx", ".vsx", ".vtx", ".vssx", ".vstx", ".vssm", ".vstm", ".vsdm", ".vssm", ".vss"},
		"media":   {".mp3", ".mp4", ".avi", ".mov", ".rmvb", ".rm", ".flv", ".mp4", ".3gp", ".mpeg", ".mpg", ".dat", ".mkv", ".wmv", ".asf", ".asx", ".vob", ".wav", ".wma", ".cda", ".flac", ".m4a", ".mid", ".mka", ".mp2", ".mpa", ".mpc", ".ape", ".ofr", ".ogg", ".ra", ".wv", ".tta", ".ac3", ".dts", ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".ico", ".webp", ".tif", ".tiff", ".pcx", ".tga", ".exif", ".fpx", ".svg", ".psd", ".cdr", ".pcd", ".dxf", ".ufo", ".eps", ".ai", ".raw", ".WMF"},
	},
}

// UserConfigs 配置文件
var UserConfigs configs

// InitConfigs 初始化配置文件
func InitConfigs() {
	//UserConfigs = defaultConfigs
	// 检查配置文件是否存在
	if _, err := os.Stat(initialize.RuntimeInfo.ConfigPath); os.IsNotExist(err) {
		// 初始化配置文件
		UserConfigs = defaultConfigs
		saveConfigs()
	} else {
		// 存在则加载
		loadConfigs()
	}
}

// 加载配置文件
func loadConfigs() {
	// 加载配置文件
	data, err := os.ReadFile(initialize.RuntimeInfo.ConfigPath)
	if err != nil {
		fmt.Println("lsh：config file load failed.")
		panic(err)
	}
	err = yaml.Unmarshal(data, &UserConfigs)
	if err != nil {
		fmt.Println("lsh：config file load failed.")
		panic(err)
	}
}

// 保存配置文件
func saveConfigs() {
	// 保存配置文件
	data, err := yaml.Marshal(UserConfigs)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(initialize.RuntimeInfo.ConfigPath, data, 0644)
	if err != nil {
		panic(err)
	}
}
