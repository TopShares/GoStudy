WebSocket
由 Http协议upgade而来
client  <-->  server


response headers:
c: upgrade       升级到WebSocket协议
s: switching     同意客户端转换WebSocket
c: message       双向通信
s: message       同上




NodeJS
    单线程，推送性能有限
C/C++
    TCP通讯，WebSocket协议实现成本高
GO!
    多线程，基于协程模型并发
    WebSocket标准库，无需造轮子
