package dlog


const (
	LOG_MAX_LINE	= 10000
	LOG_MAX_BYTE 	= 50*1024*1024
)

const (
	LOG_LVL_INFO = iota
	LOG_LVL_DEBUG
	LOG_LVL_WARN
	LOG_LVL_ERROR
	LOG_LVL_FATAL
	LOG_LVL_PANIC
	LOG_LVL_MAX
)

var (
	logLvlName	= [LOG_LVL_MAX]string{
		"info",
		"debug",
		"warn",
		"error",
		"fatal",
		"panic",
	}
)