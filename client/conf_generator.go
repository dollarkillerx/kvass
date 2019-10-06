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
)

type confGenerator struct {
}

func (c *confGenerator) Run(opt *Option) error {
	byt, e := easyutils.Base64Decode(temp.ConfTemp)
	if e != nil {
		log.Fatal(e)
	}

	e = WiteGen("config/config.go", string(byt), &IMap{"Time": easyutils.TimeGetNowTimeStr()})
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
