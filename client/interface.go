/**
 * @Author: DollarKiller
 * @Description: interface func
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:30 2019-10-05
 */
package main

type Generator interface {
	Run(opt *Option) error
}
