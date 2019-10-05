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
	Name string `json:"name"`
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
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, Conf)
	if e != nil {
		panic(e.Error())
	}
}

var conf = `
# Cli Config
app:
  - name: ""
`
