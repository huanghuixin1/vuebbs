package main

import "gopkg.in/baa.v1"
import "./router"
import (
	"os"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	//"./db"
	"fmt"
)

func main() {
	//db.TestQuery()
	b := baa.Default()
	router.Init(b)
	router.Router(b)


	config := b.GetDI("config").(map[string]string)
	fmt.Println(config)
	b.Run(":" + config["httpport"])
}

func test() {
	execDirAbsPath, _ := os.Getwd()
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)

	execFileRelativePath, _ := exec.LookPath(os.Args[0])
	log.Println("执行程序与命令执行目录的相对路径　　　　:", execFileRelativePath)

	execDirRelativePath, _ := path.Split(execFileRelativePath)
	log.Println("执行程序所在目录与命令执行目录的相对路径:", execDirRelativePath)

	execFileAbsPath, _ := filepath.Abs(execFileRelativePath)
	log.Println("执行程序的绝对路径　　　　　　　　　　　:", execFileAbsPath)

	execDirAbsPath, _ = filepath.Abs(execDirRelativePath)
	log.Println("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)

	os.Chdir(execDirRelativePath) //进入目录
	enteredDirAbsPath, _ := os.Getwd()
	log.Println("所进入目录的绝对路径　　　　　　　　　　:", enteredDirAbsPath)
}