/**
 * @Author: DollarKiller
 * @Description: main generator
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 12:12 2019-10-06
 */
package main

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/kvass/client/temp"
	"log"
)

type mainGenerator struct {
}

func (m *mainGenerator) Run(opt *Option) error {
	namePath := "main.go"
	err := WiteGen(namePath, temp.MainTemp, &IMap{
		"Time":    easyutils.TimeGetNowTimeStr(),
		"Package": opt.PackagePath,
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func init() {
	Register("mainGenerator", &mainGenerator{})
}
