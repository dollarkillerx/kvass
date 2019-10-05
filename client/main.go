/**
 * @Author: DollarKiller
 * @Description: tools cli工具
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:46 2019-10-05
 */
package main

import (
	_ "github.com/dollarkillerx/kvass/config"
	"github.com/urfave/cli"
	"log"
	"os"
)

/**
命令行参数
-p 指定包名 (如果非go mod 可以默认 系统会自己查询)
*/
func main() {
	var opt Option

	app := cli.NewApp()

	app.Name = "Kvass 微爬虫框架"
	app.Email = "dollarkiller@dollarkiller.com"
	app.Author = "DollarKiller"
	app.Copyright = "https://github.com/dollarkillerx"
	app.Version = "v0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "p",
			Value:       "",
			Usage:       "指定包名 (如果非go mod 可以默认 系统会自己查询)",
			Destination: &opt.PackagePath,
		},
	}

	app.Action = func(c *cli.Context) error {
		if opt.PackagePath == "" {
			
		}
		err := clientMag.Run(&opt)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
