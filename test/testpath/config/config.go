/**
 * @Author: DollarKiller
 * @Description: Kvass 自动生成
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 1570335735
 */
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type myconf struct {
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}
	Pgsql struct {
		Dsn     string        `yaml:"dsn"`
		MaxIdle int           `yaml:"max_idle"`
		MaxOpen int           `yaml:"max_open"`
		TimeOut time.Duration `yaml:"time_out"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	MyConfig *myconf
)

func init() {
	MyConfig = &myconf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}
}
