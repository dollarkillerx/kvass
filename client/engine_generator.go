/**
 * @Author: DollarKiller
 * @Description: engine generator
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:44 2019-10-06
 */
package main

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/kvass/client/temp"
	"github.com/dollarkillerx/kvass/config"
	"github.com/dollarkillerx/kvass/utils"
)

type engineGenerator struct {
}

func (e *engineGenerator) Run(opt *Option) error {
	filename := "fetcher/engine.go"
	data := make([]*config.Logic, 0)
	for _, k := range config.Conf.App {
		k.UpName = utils.Capitalize(k.Name)
		data = append(data, k)
	}
	err := WiteGen(filename, temp.EngineTemp, &IMap{
		"Time":    easyutils.TimeGetNowTimeStr(),
		"Maps":    data,
		"Package": opt.PackagePath,
		"EndNum":  len(data) - 1,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func init() {
	Register("engineGenerator", &engineGenerator{})
}
