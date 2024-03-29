## protobuf和grpc学习笔记
___

### 简介
记录在学习protobuf和grpc学习中的一些要点，不会太详细只是记录一些关键点，主要学习过程是参考《Go语言编程之旅》相关章节，在学习本书的时候发现相关第三方API版本更新，没有使用旧API,阅读相关文档解决部分冲突。

### grpc调用模型
1. 客户端(grpc Stub)在程序中调用某方法，发起RPC调用。
2. 对请求消息使用Protobuf进行对象序列化压缩(IDL).
3. 服务端(grpc Server)接收请求后，解码请求，进行业务逻辑处理并返回。
4. 对响应结果使用Protobuf进行对象序列化压缩(IDL).
5. 客户端接收服务端响应后，解码请求体。回调被调用的A方法，唤醒正在等待响应(阻塞)的客户端调用并返回响应结果。

### grpc优点
1. 性能好
2. 代码生成
3. 流传输
4. 超时和取消

### grpc缺点
1. 可读性差
2. 不支持浏览器调用
3. 外部组支持较差

### gprc安装
1. protoc 安装
2. proto-gen-go 安装

### 