# A Generic Go Logging Interface

Other language: 
### **[中文](README.zh.md)**
 
This is a generic Go logging interface. The purpose is to support different logging libraries and avoid to lock-in into specific logging implementation. It doesn't necessarily include all functionality for logging service, but only the ones I need. You can easily expand it to add new ones. Currently, it has implementation for [ZAP](https://github.com/uber-go/zap) and [Logrus](https://github.com/sirupsen/logrus). You can add implementation for other Messaging service like ["glog"](https://github.com/golang/glog).

For examples on how to use it in real project, please take a look at ["Order"](https://github.com/jfeng45/order) or ["Payment"](https://github.com/jfeng45/payment)

### Download Code

```
go get github.com/jfeng45/glogger
```

### License

[MIT](LICENSE.txt) License


