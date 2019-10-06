/**
 * @Author: DollarKiller
 * @Description: sync map test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:33 2019-10-05
 */
package test

import (
	"fmt"
	"github.com/dollarkillerx/easyutils"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"text/template"
)

func TestMap(t *testing.T) {
	i := sync.Map{}

	actual, loaded := i.LoadOrStore("1", "1")
	log.Println(loaded)
	log.Println(actual)

	log.Println("=========")
	actual, loaded = i.LoadOrStore("1", "2")
	log.Println(loaded)
	log.Println(actual)
}

func TestBase(t *testing.T) {
	data := `LyoqCiAqIEBBdXRob3I6IERvbGxhcktpbGxlcgogKiBARGVzY3JpcHRpb246IEt2YXNzIOiHquWKqOeUn+aIkAogKiBAR2l0aHViOiBodHRwczovL2dpdGh1Yi5jb20vZG9sbGFya2lsbGVyeAogKiBARGF0ZTogQ3JlYXRlIGluIHt7LlRpbWV9fQogKi8KcGFja2FnZSBjb25maWcKCmltcG9ydCAoCgkiZ29wa2cuaW4veWFtbC52MiIKCSJpby9pb3V0aWwiCgkidGltZSIKKQoKdHlwZSBteWNvbmYgc3RydWN0IHsKCU15c3FsIHN0cnVjdCB7CgkJRHNuICAgc3RyaW5nIGB5YW1sOiJkc24iYAoJCUNhY2hlIGJvb2wgICBgeWFtbDoiY2FjaGUiYAoJfQoJUGdzcWwgc3RydWN0IHsKCQlEc24gICAgIHN0cmluZyAgICAgICAgYHlhbWw6ImRzbiJgCgkJTWF4SWRsZSBpbnQgICAgICAgICAgIGB5YW1sOiJtYXhfaWRsZSJgCgkJTWF4T3BlbiBpbnQgICAgICAgICAgIGB5YW1sOiJtYXhfb3BlbiJgCgkJVGltZU91dCB0aW1lLkR1cmF0aW9uIGB5YW1sOiJ0aW1lX291dCJgCgl9CglSZWRpcyBzdHJ1Y3QgewoJCU1heGlkbGUgICAgIGludCAgICAgICAgICAgYHlhbWw6Im1heGlkbGUiYAoJCU1heEFjdGl2ZSAgIGludCAgICAgICAgICAgYHlhbWw6Im1heF9hY3RpdmUiYAoJCUlkbGVUaW1lb3V0IHRpbWUuRHVyYXRpb24gYHlhbWw6ImlkbGVfdGltZW91dCJgCgkJUG9ydCAgICAgICAgc3RyaW5nICAgICAgICBgeWFtbDoicG9ydCJgCgl9Cn0KCnZhciAoCglNeUNvbmZpZyAqbXljb25mCikKCmZ1bmMgaW5pdCgpIHsKCU15Q29uZmlnID0gJm15Y29uZnt9CgoJYnl0ZXMsIGUgOj0gaW91dGlsLlJlYWRGaWxlKCIuL2NvbmZpZy55bWwiKQoJaWYgZSAhPSBuaWwgewoJCXBhbmljKGUuRXJyb3IoKSkKCX0KCgllID0geWFtbC5Vbm1hcnNoYWwoYnl0ZXMsIE15Q29uZmlnKQoJaWYgZSAhPSBuaWwgewoJCXBhbmljKGUuRXJyb3IoKSkKCX0KfQo=`
	bytes, e := easyutils.Base64Decode(data)
	if e != nil {
		log.Fatal(e)
	}

	log.Println(string(bytes))
}

func TestHtml(t *testing.T) {
	data := `LyoqCiAqIEBBdXRob3I6IERvbGxhcktpbGxlcgogKiBARGVzY3JpcHRpb246IEt2YXNzIOiHquWKqOeUn+aIkAogKiBAR2l0aHViOiBodHRwczovL2dpdGh1Yi5jb20vZG9sbGFya2lsbGVyeAogKiBARGF0ZTogQ3JlYXRlIGluIHt7LlRpbWV9fQogKi8KcGFja2FnZSBjb25maWcKCmltcG9ydCAoCgkiZ29wa2cuaW4veWFtbC52MiIKCSJpby9pb3V0aWwiCgkidGltZSIKKQoKdHlwZSBteWNvbmYgc3RydWN0IHsKCU15c3FsIHN0cnVjdCB7CgkJRHNuICAgc3RyaW5nIGB5YW1sOiJkc24iYAoJCUNhY2hlIGJvb2wgICBgeWFtbDoiY2FjaGUiYAoJfQoJUGdzcWwgc3RydWN0IHsKCQlEc24gICAgIHN0cmluZyAgICAgICAgYHlhbWw6ImRzbiJgCgkJTWF4SWRsZSBpbnQgICAgICAgICAgIGB5YW1sOiJtYXhfaWRsZSJgCgkJTWF4T3BlbiBpbnQgICAgICAgICAgIGB5YW1sOiJtYXhfb3BlbiJgCgkJVGltZU91dCB0aW1lLkR1cmF0aW9uIGB5YW1sOiJ0aW1lX291dCJgCgl9CglSZWRpcyBzdHJ1Y3QgewoJCU1heGlkbGUgICAgIGludCAgICAgICAgICAgYHlhbWw6Im1heGlkbGUiYAoJCU1heEFjdGl2ZSAgIGludCAgICAgICAgICAgYHlhbWw6Im1heF9hY3RpdmUiYAoJCUlkbGVUaW1lb3V0IHRpbWUuRHVyYXRpb24gYHlhbWw6ImlkbGVfdGltZW91dCJgCgkJUG9ydCAgICAgICAgc3RyaW5nICAgICAgICBgeWFtbDoicG9ydCJgCgl9Cn0KCnZhciAoCglNeUNvbmZpZyAqbXljb25mCikKCmZ1bmMgaW5pdCgpIHsKCU15Q29uZmlnID0gJm15Y29uZnt9CgoJYnl0ZXMsIGUgOj0gaW91dGlsLlJlYWRGaWxlKCIuL2NvbmZpZy55bWwiKQoJaWYgZSAhPSBuaWwgewoJCXBhbmljKGUuRXJyb3IoKSkKCX0KCgllID0geWFtbC5Vbm1hcnNoYWwoYnl0ZXMsIE15Q29uZmlnKQoJaWYgZSAhPSBuaWwgewoJCXBhbmljKGUuRXJyb3IoKSkKCX0KfQo=`
	byt, e := easyutils.Base64Decode(data)
	if e != nil {
		log.Fatal(e)
	}
	file, e := os.OpenFile("h1.html", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 00666)
	if e != nil {
		log.Fatal(e)
	}
	html := string(byt)

	i := template.New("test")
	i.Parse(html)
	e = i.Execute(file, map[string]interface{}{
		"Time": "hellll",
	})
	if e != nil {
		log.Fatal(e)
	}
}

func TestPath(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	log.Println(goPath)
	exeFilePath, err := filepath.Abs(os.Args[0])

	if err == nil {
		log.Println(exeFilePath)
	} else {
		log.Fatal(err)
	}
}

func TestPcc(t *testing.T) {
	ti := "ffc"
	title := capitalize(ti)
	log.Println(title)
}

func capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func TestTemp(t *testing.T) {
	data := []string{
		"Hello",
		"TTc",
	}
	file, e := os.OpenFile("test.html", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 00755)
	if e != nil {
		log.Fatal(e)
	}
	i := template.New("test")
	i.Funcs(template.FuncMap{"tep": testHtml})
	i.Parse(testhtml)
	e = i.Execute(file, map[string]interface{}{
		"data": data,
	})
	if e != nil {
		log.Fatal(e)
	}
}

func testHtml(num int, data []string) string {
	return data[num+1]
}

var testhtml = `
	// this is data
	{{range $k,$v := .data}}
		{{if eq $k 0}}
			{{tep $k $.data}}
		{{end}}		
	{{end}}
	// end
`

var c = `
{{if eq $k 0}}
{{$k $.data | tep}}
{{end}}
`
