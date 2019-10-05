/**
 * @Author: DollarKiller
 * @Description: client tools mag 综合管理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:12 2019-10-05
 */
package main

import (
	"log"
	"os"
	"sync"
	"text/template"
)

type IMap map[string]interface{}

var (
	clientMag = &generatorMgr{
		genMap: sync.Map{},
	}
)

type generatorMgr struct {
	genMap sync.Map
}

func (c *generatorMgr) Run(opt *Option) error {
	// 创建目录
	c.createDirectory()
	return nil
}

var dirPath = []string{
	"config",
	"datamodels",
	"datasources/mysql_con",
	"datasources/pgsql_con",
	"defs",
	"fetcher",
	"reptile",
	"test",
	"utils",
}

func (c *generatorMgr) createDirectory() {
	for _, path := range dirPath {
		err := os.MkdirAll(path, 00755)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("目录初始化完毕")
}

func Register(name string, generator Generator) {
	clientMag.genMap.LoadOrStore(name, generator)
}

// 写入
func WiteGen(file *os.File, data string, metaData *IMap) error {
	i := template.New("config")
	i.Parse(string(data))
	return i.Execute(file, metaData)
}
