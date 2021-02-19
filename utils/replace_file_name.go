package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

/**
批量修改文件夹内 文件名称
*/
func main() {
	// 解析命令行参数
	var dir string
	flag.StringVar(&dir, "d", "", "文件夹名称")
	var pattern string
	flag.StringVar(&pattern, "p", "", "name pattern, eg. newname%d")

	flag.Parse()
	if dir == "" || pattern == "" {
		flag.Usage()
		return
	}

	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	// 遍历文件路径，修改文件名
	for _, path := range paths {
		newPath := strings.Replace(path, pattern, "", -1)
		//newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf(pattern, i+1)+filepath.Ext(path))
		os.Rename(path, newPath)
	}
}
