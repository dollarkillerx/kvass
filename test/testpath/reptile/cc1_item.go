
/**
 * @Author: DollarKiller
 * @Description: 由 Kvass cli生成
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 1570335735
 */
package reptile

import (
	"github.com/dollarkillerx/easyutils/clog"
	"sync"
)

// parser home
type Cc1 struct {
}

func (p *Cc1) ParserItem(ch1 chan interface{}, ch2 chan interface{}) {
	numch := make(chan int, 2)
	sy := sync.WaitGroup{}
cc:
	for {
		select {
		case ur, ok := <-ch1:
			if ok {
				// 开启多协程
				numch <- 1
				sy.Add(1)
				go func(ur interface{}) {
					defer func() {
						<-numch
						sy.Done()
					}()
					p.logic(ur, ch2)

				}(ur)

			} else {
				sy.Wait()
				clog.Println("第1阶段完毕")
				close(ch2)
				break cc
			}
		}
	}
}

func (p *Cc1) logic(data interface{}, ch chan interface{}) {
	ch <- "test"
}
