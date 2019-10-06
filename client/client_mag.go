/**
 * @Author: DollarKiller
 * @Description: client tools mag 综合管理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:12 2019-10-05
 */
package main

import (
	"github.com/dollarkillerx/kvass/config"
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
	c.genMap.Range(func(key, value interface{}) bool {
		generator := value.(Generator)
		generator.Run(opt)
		return true
	})
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
func WiteGen(filename string, data string, metaData *IMap) error {
	file, e := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 00755)
	if e != nil {
		return e
	}
	i := template.New("config")
	i.Funcs(template.FuncMap{
		"pxy": pxy,
	})
	i.Parse(string(data))
	return i.Execute(file, metaData)
}

func pxy(num int, maps []*config.Logic) string {
	return maps[num-1].Name
}
