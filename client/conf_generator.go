/**
 * @Author: DollarKiller
 * @Description: config generator
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:37 2019-10-05
 */
package main

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/kvass/client/temp"
	"log"
	"os"
)

type confGenerator struct {
}

func (c *confGenerator) Run(opt *Option) error {
	file, e := os.OpenFile("config/config.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 00755)
	if e != nil {
		log.Fatal(e)
	}

	byt, e := easyutils.Base64Decode(temp.ConfTemp)
	if e != nil {
		log.Fatal(e)
	}

	e = WiteGen(file, string(byt), &IMap{"Time": easyutils.TimeGetNowTimeStr()})
	if e != nil {
		clog.PrintWa(e)
		log.Fatal("err 程序运行异常")
	}

	return nil
}

func init() {
	generator := confGenerator{}
	Register("config_gen", &generator)
}
