/**
 * @Author: DollarKiller
 * @Description: reptile generator 单个爬取逻辑生成
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:01 2019-10-05
 */
package main

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/kvass/client/temp"
	"github.com/dollarkillerx/kvass/config"
	"github.com/dollarkillerx/kvass/utils"
	"log"
)

type ReptileGenerator struct {
}

func (r *ReptileGenerator) Run(opt *Option) error {
	for i, item := range config.Conf.App {
		if i == 0 {
			// i == 0 为分发层
			filename := "reptile/" + item.Name + "_item.go"
			e := WiteGen(filename, temp.ReptileDistribution, &IMap{
				"Time":    easyutils.TimeGetNowTimeStr(),
				"FunName": utils.Capitalize(item.Name),
			})
			if e != nil {
				log.Fatal(e)
			}
		} else {
			// 应用层
			filename := "reptile/" + item.Name + "_item.go"
			e := WiteGen(filename, temp.ReptileItem, &IMap{
				"Time":    easyutils.TimeGetNowTimeStr(),
				"FunName": utils.Capitalize(item.Name),
				"Num":     item.Num,
				"Nu":      i,
			})
			if e != nil {
				log.Fatal(e)
			}
		}

	}
	return nil
}

func init() {
	generator := ReptileGenerator{}
	Register("ReptileGenerator", &generator)
}
