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

内部函数的简单使用:
```
package main

import (
	"fmt"
	"github.com/dzhenquan/dlog"
)

func main() {
	//当前存放日志的目录      就日志文件目录       日志文件前缀  日志文件后缀(默认log)
	log := dlog.New("./log", "./old_log", "test", "")

	logDir1 := log.GetLogDir()   				// 获取保存日志文件的目录
	fmt.Println("logDir1:", logDir1)

	log.SetLogDir("./dai")					// 设置保存日志文件的目录
	logDir2 := log.GetLogDir()
	fmt.Printf("logDir2: %s\n\n", logDir2)


	logOldDir1 := log.GetLogOldDir()			// 获取保存日志文件的旧目录
	fmt.Println("logOldDir1:", logOldDir1)

	log.SetLogOldDir("./old_dai")		// 设置保存日志文件的旧目录
	logOldDir2 := log.GetLogOldDir()
	fmt.Printf("logOldDir2: %s\n\n", logOldDir2)


	logPre1 := log.GetPrefix()					// 获取日志文件的前缀名
	fmt.Println("logPre1:", logPre1)

	log.SetPrefix("test")					// 设置日志文件的前缀名
	logPre2 := log.GetPrefix()
	fmt.Printf("logPre2: %s\n\n", logPre2)


	logSuf1 := log.GetSuffix()					// 获取日志文件的后缀名
	fmt.Println("logSuf1:", logSuf1)

	log.SetSuffix("txt")
	logSuf2 := log.GetSuffix()					// 设置日志文件的后缀名
	fmt.Printf("logSuf2: %s\n\n", logSuf2)


	logMaxLine1 := log.GetMaxLine()				// 获取日志文件的最大行数(默认10000行)
	fmt.Println("logMaxLine1:", logMaxLine1)

	log.SetMaxLine(99)
	logMaxLine2 := log.GetMaxLine()				// 设置日志文件的最大行数
	fmt.Printf("logMaxLine2: %d\n\n", logMaxLine2)


	logMaxByte1 := log.GetMaxByte()				// 获取日志文件的最大字节数(默认50M)
	fmt.Println("logMaxByte1:", logMaxByte1)

	log.SetMaxByte(100*1024*1024)			// 设置日志文件的最大字节数
	logMaxByte2 := log.GetMaxByte()
	fmt.Printf("logMaxByte2: %d\n\n", logMaxByte2)
}
```
