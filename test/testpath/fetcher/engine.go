
/**
 * @Author: DollarKiller
 * @Description: 由 Kvass cli生成
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 1570335735
 */
package fetcher

import (
	"github.com/dollarkillerx/kvass/test/testpath/reptile"
)

// 分发url
type ParserUrl interface {
	ParserUrl(chan interface{})
}

// 解析url
type ParserItem interface {
	ParserItem(chan interface{}, chan interface{})
}

type Reptile struct {
	
	ginCh chan interface{}
	
	cc1Ch chan interface{}
	
	cc2Ch chan interface{}
	
	
		
	GenerateUrl ParserUrl
		
	
		
	cc1 ParserItem
		
	
		
	cc2 ParserItem
		
	
}

// 中央控制
func ReptileEngine() {
	
		
	url := reptile.GenerateUrl{}    // url 生成器
		
	
		
	cc1 := reptile.Cc1{}
		
	
		
	cc2 := reptile.Cc2{}
		
	
	i := Reptile{
		
		ginCh: make(chan interface{}, 15),
		
		cc1Ch: make(chan interface{}, 15),
		
		cc2Ch: make(chan interface{}, 15),
		

		
			
		GenerateUrl: &url,
			
		
			
		cc1: &cc1,
			
		
			
		cc2: &cc2,
			
		
	}
	
		
	go i.GenerateUrl.ParserUrl(i.ginCh)
		
	
		
	go i.cc1.ParserItem(i.ginCh, i.cc1Ch)	
		
	
		
	go i.cc2.ParserItem(i.cc1Ch, i.cc2Ch)
	<-i.cc2Ch
		
	
}
