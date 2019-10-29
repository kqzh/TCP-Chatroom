# TCP-Chatroom
使用 Golang 搭建的网络聊天室

## 启动服务端

```go
go run serve.go
```
## 启动客户端

```go
go run client.go
```
输入昵称后进入聊天室
```go
please set your name : kqzh
```
自动获取 Usage
```go
welcome visitor, you can send message to anyone who is online too!

send message 'list' to get online members

send message 'quit' to exit

message format is 'xxx #name'

```
## 操作演示

### 查看当前在线成员
```go
list
```
输出聊天室所有成员
```go
online members : kqzh chen
```
### 发送消息给chen
```go
hello , can you hear me? #chen
```
chen 会收到来自kqzh的消息
```go
kqzh : hello , can you hear me?
```

### 退出聊天室
```go
quit
```
客户端停止运行
```go
have a good day!
```

