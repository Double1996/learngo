package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/**
批量修改文件夹内 文件名称
*/
var rMap = map[string]string{
	"格律法学院：": "",
	"-------格律法学院（微信订阅号：gelvfaxueyuan）": "",
	"-格律法学院（微信公众号gelvfaxueyuan)":        "",
	"------格律法学院（微信订阅号：gelvfaxueyuan）":  "",
}

var dirs = []string{
	"C:\\Users\\Administrator\\Downloads\\1741个合同范本",
	"C:\\Users\\Administrator\\Downloads\\格律法学院：1080个各省公布的合同范本\\格律法学院：1080个各省公布的合同范本",
}

func main() {
	// 解析命令行参数
	var dir string
	flag.StringVar(&dir, "d", "", "文件夹名称")

	flag.Parse()

	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	for _, dir := range dirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})
	}
	// 遍历文件路径，修改文件名
	i := 0
	for _, path := range paths {
		for k, v := range rMap {
			newPath := strings.Replace(path, k, v, -1)
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(i, newPath)
		}
		i += 1
		fmt.Println(i, path)

	}
	fmt.Printf("-------------sucess-----------------------")
}
