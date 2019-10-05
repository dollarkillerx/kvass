/**
 * @Author: DollarKiller
 * @Description: temp template
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:44 2019-10-05
 */
package temp

var ConfTemp = `
LyoqCiAqIEBBdXRob3I6IERvbGxhcktpbGxlcgogKiBARGVzY3JpcHRpb246IEt2YXNzIOiHquWKqOeUn+aIkAogKiBAR2l0aHViOiBodHRwczovL2dpdGh1Yi5jb20vZG9sbGFya2lsbGVyeAogKiBARGF0ZTogQ3JlYXRlIGluIHt7LlRpbWV9fQogKi8KcGFja2FnZSBjb25maWcKCmltcG9ydCAoCgkiZ29wa2cuaW4veWFtbC52MiIKCSJpby9pb3V0aWwiCgkidGltZSIKKQoKdHlwZSBteWNvbmYgc3RydWN0IHsKCU15c3FsIHN0cnVjdCB7CgkJRHNuICAgc3RyaW5nIGB5YW1sOiJkc24iYAoJCUNhY2hlIGJvb2wgICBgeWFtbDoiY2FjaGUiYAoJfQoJUGdzcWwgc3RydWN0IHsKCQlEc24gICAgIHN0cmluZyAgICAgICAgYHlhbWw6ImRzbiJgCgkJTWF4SWRsZSBpbnQgICAgICAgICAgIGB5YW1sOiJtYXhfaWRsZSJgCgkJTWF4T3BlbiBpbnQgICAgICAgICAgIGB5YW1sOiJtYXhfb3BlbiJgCgkJVGltZU91dCB0aW1lLkR1cmF0aW9uIGB5YW1sOiJ0aW1lX291dCJgCgl9CglSZWRpcyBzdHJ1Y3QgewoJCU1heGlkbGUgICAgIGludCAgICAgICAgICAgYHlhbWw6Im1heGlkbGUiYAoJCU1heEFjdGl2ZSAgIGludCAgICAgICAgICAgYHlhbWw6Im1heF9hY3RpdmUiYAoJCUlkbGVUaW1lb3V0IHRpbWUuRHVyYXRpb24gYHlhbWw6ImlkbGVfdGltZW91dCJgCgkJUG9ydCAgICAgICAgc3RyaW5nICAgICAgICBgeWFtbDoicG9ydCJgCgl9Cn0KCnZhciAoCglNeUNvbmZpZyAqbXljb25mCikKCmZ1bmMgaW5pdCgpIHsKCU15Q29uZmlnID0gJm15Y29uZnt9CgoJYnl0ZXMsIGUgOj0gaW91dGlsLlJlYWRGaWxlKCIuL2NvbmZpZy55bWwiKQoJaWYgZSAhPSBuaWwgewoJCXBhbmljKGUuRXJyb3IoKSkKCX0KCgllID0geWFtbC5Vbm1hcnNoYWwoYnl0ZXMsIE15Q29uZmlnKQoJaWYgZSAhPSBuaWwgewoJCXBhbmljKGUuRXJyb3IoKSkKCX0KfQo=
`

var ConfigTemp = `
# Kvass 自动生成 通用配置文件

app:
  host: "0.0.0.0:8083"
  log_level: "debug"
  debug: true
  max_request: 1000

pgsql:
  dsn: "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword"
  max_idle: 10 # 最大闲置链接数
  max_open: 100 # 最大链接数
  time_out: 4 # 超时 Hour

redis:
  maxidle: 50 # 连接池中最大空闲连接数量
  max_active: 30 # 最大连接数量
  idle_timeout: 300 # 最大超时时间(s)
  port: "127.0.0.1:6379"
`
