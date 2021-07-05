package biz

import (
	"io/ioutil"
)

type File struct {
}

// ReadFromDisk 从磁盘读取文件,返回字节数组
func ReadFromDisk(filePath string) []byte {
	bytes, _ := ioutil.ReadFile(filePath)
	return bytes
}
