package log

import (
	"fmt"
	baidulog "icode.baidu.com/baidu/go-lib/log"
	"icode.baidu.com/baidu/go-lib/log/log4go"
	"icode.baidu.com/baidu/goodcoder/gongyitong/constants"

	"sync"
	"time"
)

var once sync.Once
var lastSendTime time.Time

//true的情况下会额外使用fmt.Printf打印到strout,仅验功能则不需要额外打印
var IsDebugMode = false
var Env = "prod"

func setup() error {
	log4go.SetLogFormat(log4go.FORMAT_JSON)
	baidulog.Init("test", "INFO", "./log", true, "midnight", 5)
	return nil
}

func Debug(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	baidulog.Logger.Debug(format, v...)
}

// Legacy functions
func Critical(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	baidulog.Logger.Critical(format, v...)
}

func Err(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	baidulog.Logger.Error(format, v...)

}

func Info(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	baidulog.Logger.Info(format, v...)

}

func Warning(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	baidulog.Logger.Warn(format, v...)

}


func Duration(t time.Time, ctx string) {
	Debug("%.4fs for %s", time.Now().Sub(t).Seconds(), ctx)
}

func SetDebugMode(is bool) {
	IsDebugMode = is
	Info("debug mode set to %v", is)
}

func SetEnv(env string) {
	Env = env
	Info("env set to %s", env)
}

