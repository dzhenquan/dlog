dlog 是一个用于保存日志文件，简单而又多功能的工具,由于系统自带的log用这不是太舒服，  
满足不了我的某些需求。所以，自己造了一个简单的轮子，用于保存日志文件。  

首先将源代码下载到gopath环境
```
[root@localhost ~]# go get github.com/dzhenquan/dlog
```
或者直接将源代码下载放到自己的项目中使用
```
git clone https://github.com/dzhenquan/dlog.git
```

简单使用方法如下：
```
package main

import (
	"fmt"
	"time"
	"github.com/dzhenquan/dlog"
)

func testPrint(log *dlog.DLogger) {
	for i := 0; i < 100; i ++ {
		log.Print("i = ", i, "\n")
		log.Printf("i: [%d]\n", i)
		log.Println("i = ", i)
		time.Sleep(200*time.Millisecond)
	}
}

func testDebug(log *dlog.DLogger) {
	for i := 0; i < 100; i ++ {
		log.Debug("i = ", i, "\n")
		log.Debugf("i: [%d]\n", i)
		log.Debugln("i = ", i)
		time.Sleep(200*time.Millisecond)
	}
}

func testWarn(log *dlog.DLogger) {
	for i := 0; i < 100; i ++ {
		log.Warn("i = ", i, "\n")
		log.Warnf("i: [%d]\n", i)
		log.Warnln("i = ", i)
		time.Sleep(200*time.Millisecond)
	}
}

func testError(log *dlog.DLogger) {
	for i := 0; i < 100; i ++ {
		log.Error("i = ", i, "\n")
		log.Errorf("i: [%d]\n", i)
		log.Errorln("i = ", i)
		time.Sleep(200*time.Millisecond)
	}
}

func main() {
	//当前存放日志的目录      就日志文件目录       日志文件前缀  日志文件后缀(默认log)
	log := dlog.New("./log", "./old_log", "test", "")  //首先创建一个日志对象

	fmt.Println("------------- Start Test ---------------")

	go testPrint(log)

	go testDebug(log)

	go testWarn(log)

	go testError(log)

	for {
		time.Sleep(20*time.Second)
	}

	fmt.Println("------------- End Test ---------------")
}
```
