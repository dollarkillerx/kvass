/**
 * @Author: DollarKiller
 * @Description: 特殊的测试
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:32 2019-10-05
 */
package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	goPath := os.Getenv("GOPATH")
	exeFilePath, err := filepath.Abs(os.Args[0])

	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		exeFilePath = strings.Replace(exeFilePath, "\\", "/", -1)
	}

	goPath = goPath + "/src/"
	data := strings.Replace(exeFilePath, goPath, "", -1)
	log.Println(data)
}
