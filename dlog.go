package dlog

import (
	"syscall"
	"sync"
	"os"
	"time"
	"fmt"
	"runtime"
)

type DLogger struct {
	file 		*os.File
	logDir 		string			// 当前日志文件存放目录
	logOldDir	string			// 旧日志文件存放目录
	logPre 		string			// 日志文件前缀
	logSuf 		string			// 日志文件后缀
	dlogLvl		[LOG_LVL_MAX]DLogLevel
	mutex 		*sync.Mutex
}

type DLogLevel struct {
	logLine		int64			// 日志文件行数
	logByte 	int64			// 日志文件字节数
	logCount 	int64			// 日志文件数
	logMaxLine 	int64			// 日志文件最大行数
	logMaxByte	int64			// 日志文件最大字节数
}

func New(logDir, logOldDir,logPre,logSuf string) *DLogger {
	if len(logDir) <= 0 {
		logDir = "."
	}

	if len(logOldDir) <= 0 {
		logOldDir = "."
	}

	if len(logPre) <= 0 {
		logPre = "program"
	}

	if len(logSuf) <= 0 {
		logSuf = "log"
	}

	logMaxLine := LOG_MAX_LINE

	logMaxByte := LOG_MAX_BYTE

	if err := syscall.Access(logDir, syscall.O_RDWR); err != nil {
		return nil
	}

	if err := syscall.Access(logOldDir, syscall.O_RDWR); err != nil {
		return nil
	}

	dlogger := &DLogger{
		logDir:logDir,
		logOldDir:logOldDir,
		logPre:logPre,
		logSuf:logSuf,
		mutex: &sync.Mutex{},
	}

	for i := 0; i < LOG_LVL_MAX; i++ {
		dlogger.dlogLvl[i].logLine = 0
		dlogger.dlogLvl[i].logByte = 0
		dlogger.dlogLvl[i].logCount = 0
		dlogger.dlogLvl[i].logMaxLine = int64(logMaxLine)
		dlogger.dlogLvl[i].logMaxByte = int64(logMaxByte)
	}

	return dlogger
}

func (dlog *DLogger) Print(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_INFO, fmt.Sprint(v...))
}

func (dlog *DLogger) Printf(format string, v ...interface{}) {
	dlog.outPut(2, LOG_LVL_INFO, fmt.Sprintf(format, v ...))
}

func (dlog *DLogger) Println(v ...interface{}) {

	dlog.outPut(2, LOG_LVL_INFO, fmt.Sprintln(v...))
}

func (dlog *DLogger) Debug(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_DEBUG, fmt.Sprint(v...))
}

func (dlog *DLogger) Debugf(format string, v ...interface{}) {
	dlog.outPut(2, LOG_LVL_DEBUG, fmt.Sprintf(format, v ...))
}

func (dlog *DLogger) Debugln(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_DEBUG, fmt.Sprintln(v...))
}

func (dlog *DLogger) Warn(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_WARN, fmt.Sprint(v...))
}

func (dlog *DLogger) Warnf(format string, v ...interface{}) {
	dlog.outPut(2, LOG_LVL_WARN, fmt.Sprintf(format, v ...))
}

func (dlog *DLogger) Warnln(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_WARN, fmt.Sprintln(v...))
}

func (dlog *DLogger) Error(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_ERROR, fmt.Sprint(v...))
}

func (dlog *DLogger) Errorf(format string, v ...interface{}) {
	dlog.outPut(2, LOG_LVL_ERROR, fmt.Sprintf(format, v ...))
}

func (dlog *DLogger) Errorln(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_ERROR, fmt.Sprintln(v...))
}

func (dlog *DLogger) Fatal(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_FATAL, fmt.Sprint(v...))
	os.Exit(1)
}

func (dlog *DLogger) Fatalf(format string, v ...interface{}) {
	dlog.outPut(2, LOG_LVL_FATAL, fmt.Sprintf(format, v ...))
	os.Exit(1)
}

func (dlog *DLogger) Fatalln(v ...interface{}) {
	dlog.outPut(2, LOG_LVL_FATAL, fmt.Sprintln(v...))
	os.Exit(1)
}

func (dlog *DLogger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	dlog.outPut(2, LOG_LVL_PANIC, s)
	panic(s)
}

// 设置保存日志文件的目录
func (dlog *DLogger) SetLogDir(logDir string) {
	if len(dlog.logDir) > 0 {
		dlog.logDir = logDir
	} else {
		dlog.logDir = "."
	}
}

// 设置保存日志文件的旧目录
func (dlog *DLogger) SetLogOldDir(logOldDir string) {
	if len(dlog.logOldDir) > 0 {
		dlog.logOldDir = logOldDir
	} else {
		dlog.logOldDir = "."
	}
}

// 设置获取保存日志文件的旧目录
func (dlog *DLogger) SetPrefix(prefix string) {
	dlog.mutex.Lock()
	defer dlog.mutex.Unlock()
	dlog.logPre = prefix
}

// 设置日志文件的后缀名
func (dlog *DLogger) SetSuffix(suffix string) {
	dlog.mutex.Lock()
	defer dlog.mutex.Unlock()
	dlog.logSuf = suffix
}

// 设置日志文件的最大行数
func (dlog *DLogger) SetMaxLine(line int64) {
	for i := 0; i < LOG_LVL_MAX; i++ {
		dlog.dlogLvl[i].logMaxLine = line
	}
}

// 设置日志文件的最大字节数
func (dlog *DLogger) SetMaxByte(byte int64) {
	for i := 0; i < LOG_LVL_MAX; i++ {
		dlog.dlogLvl[i].logMaxByte = byte
	}
}

// 获取保存日志文件的目录
func (dlog *DLogger) GetLogDir() string {
	return dlog.logDir
}

// 获取保存日志文件的旧目录
func (dlog *DLogger) GetLogOldDir() string {
	return dlog.logOldDir
}

// 获取日志文件的前缀名
func (dlog *DLogger) GetPrefix() string {
	return dlog.logPre
}

// 获取日志文件的后缀名
func (dlog *DLogger) GetSuffix() string {
	return dlog.logSuf
}

// 获取日志文件的最大行数
func (dlog *DLogger) GetMaxLine() int64 {
	return dlog.dlogLvl[0].logMaxLine
}

// 获取日志文件的最大字节数
func (dlog *DLogger) GetMaxByte() int64 {
	return dlog.dlogLvl[0].logMaxByte
}

// 输出到日志文件
func (dlog *DLogger) outPut(calldepth int, loglvl int, content string) error {
	dlog.mutex.Lock()
	defer dlog.mutex.Unlock()

	_, filename, line, ok := runtime.Caller(calldepth)
	if !ok {
		return nil
	}
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	logBuffer := fmt.Sprintf("[%s][%s:%d]", nowTime, dlog.formatFileName(filename), line)

	logBuffer += content

	dlog.logSwitch(loglvl)

	wsize, err := dlog.file.WriteString(logBuffer)
	if err != nil {
		return err
	}

	dlog.dlogLvl[loglvl].logLine++
	dlog.dlogLvl[loglvl].logByte += int64(wsize)
	dlog.file.Close()

	return nil
}

// 日志文件的交换
func (dlog *DLogger) logSwitch(loglvl int) error {
	currentFileName := dlog.getLogCurrentPathName(loglvl)

	if ((dlog.dlogLvl[loglvl].logLine > 0 &&
		dlog.dlogLvl[loglvl].logLine > dlog.dlogLvl[loglvl].logMaxLine) ||
		(dlog.dlogLvl[loglvl].logByte > 0 &&
		dlog.dlogLvl[loglvl].logByte > dlog.dlogLvl[loglvl].logMaxByte)) {

		switchFileName := dlog.getLogSwitchPathName(loglvl)

		os.Rename(currentFileName, switchFileName)
		dlog.dlogLvl[loglvl].logLine = 0
		dlog.dlogLvl[loglvl].logByte = 0
		dlog.dlogLvl[loglvl].logCount++
	}

	file, err := os.OpenFile(currentFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	dlog.file = file

	return nil
}

// 获取当前文件名
func (dlog *DLogger) getLogCurrentPathName(loglvl int) string {
	filename := fmt.Sprintf("%s/%s_%s_%s.%s",
		dlog.logDir, dlog.logPre, logLvlName[loglvl], "current", dlog.logSuf)

	return filename
}

// 获取交换文件的文件名
func (dlog *DLogger) getLogSwitchPathName(loglvl int) string {
	nowTime := time.Now().Format("20060102")

	filename := fmt.Sprintf("%s/%s_%s_%s_%03d.%s",
		dlog.logOldDir, dlog.logPre, logLvlName[loglvl],nowTime, dlog.dlogLvl[loglvl].logCount, dlog.logSuf)

	return filename
}

// 根据文件全路径获取文件名
func (dlog *DLogger) formatFileName (filePath string) string {
	filename := filePath
	for i := len(filePath)-1; i > 0; i-- {
		if filePath[i] == '/' {
			filename = filePath[i+1:]
			break
		}
	}
	filePath = filename
	return filePath
}