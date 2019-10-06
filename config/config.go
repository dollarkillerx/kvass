/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:55 2019-10-05
 */
package config

import (
	"github.com/dollarkillerx/easyutils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Logic struct {
	Name   string `json:"name"`
	UpName string `json:"-"`
	Num    string `json:"num"`
}

type myconf struct {
	App []*Logic `yaml:"app"`
}

var (
	Conf *myconf
)

func init() {
	Conf = &myconf{}

	b, i := easyutils.PathExists("cli_config.yml")
	if i != nil || b == false {
		// 如果不存在就创建
		ioutil.WriteFile("cli_config.yml", []byte(conf), 00666)
		log.Fatalf("请填写配置文件然后再次运行")
	}

	bytes, e := ioutil.ReadFile("cli_config.yml")
	if e != nil {
		log.Fatal(e)
	}

	e = yaml.Unmarshal(bytes, Conf)
	if e != nil {
		log.Fatal(e)
	}
}

var conf = `
# Kvass 生成配置文件
app:
  -
    name: "gin"    # 第一层为分发层
    num: 1          # 开启线程数
  -
    name: "cc1"    # 普通层  复数
    num: 2
  -
    name: "cc2"
    num: 3
`
