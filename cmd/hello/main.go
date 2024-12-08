package main

import (
	"flag"
	"os"

	"github.com/codeExpert666/hello"
)

var name string

// init() 函数在程序启动时自动执行，且在 main() 函数之前执行
// 用途：1、程序初始化配置 2、注册数据库驱动 3、初始化全局变量 4、建立程序运行必需的资源连接
func init() {
	flag.StringVar(&name, "name", "world", "name to say hello")
}

func main() {
	flag.Parse()
	msg := hello.Hello(name) + "\n"
	_, err := os.Stdout.WriteString(msg)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
