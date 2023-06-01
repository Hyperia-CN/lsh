/*
@Time : 2023/5/30 14:46
@Author : Hyperia
@File : metadata
@Software: GoLand
@Blog : https://blog.hyperia.cn
*/

package lib

import "github.com/pkg/xattr"

// SetAttr 修改文件元数据、增加指定的元数据字段。
func SetAttr(filename string, name string, value string) error {
	return xattr.Set(filename, name, []byte(value))
}

// GetAttr 获取指元数据中指定字段的值。
func GetAttr(filename, name string) (string, error) {
	value, err := xattr.Get(filename, name)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// DelAttr 从指定文件中移除命名属性。
func DelAttr(filename string, name string) error {
	return xattr.Remove(filename, name)
}
