## context

> [深度解密Go语言之context](https://zhuanlan.zhihu.com/p/68792989)

> [腾讯云开发者手册 context](https://cloud.tencent.com/developer/section/1140703)


### 什么是 context
GO 1.7标准库引入context,中文译作"上下文"，准确说它是goroutine的上下文，包含goroutine 的运行状态，环境，现场等信息。

context 主要用来在goroutine 之间传递上下文信息，包括: 取消信号、超时时间、截止时间、k-v等。

随着context 包的引入，标准库中很多接口因此加上context 参数,例如 database/sql包。context 几乎成为了并发控制和超时控制的标准做法.

> context.Context 类型的值可以协调多个goroutine 中代码执行 "取消"操作，例如：取消一个HTTP请求的执行。


### 为什么有 context
Go 通常用来写后台服务，通常只需要几行代码，就可以搭建一个http server.

在 Go 的 server 里，通常每来一个请求都会启动若干个 goroutine 同时工作: 有些去数据库拿数据，有些调用下游接口获取相关数据......


