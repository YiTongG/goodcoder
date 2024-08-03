package log

import (
	"fmt"
	"log"
	"os"

	"time"
)

// EnvDevbox 环境
var EnvDevbox = "dev"

// DateTimeFmt3 时间格式
var DateTimeFmt3 = "2015-02-25 15:04:05.000"

// IsDebugMode true的情况下会额外使用fmt.Printf打印到strout,仅验功能则不需要额外打印
var IsDebugMode = false

// Env 环境名
var Env = "prod"

// SetDebugMode 设置为debug模式，调试时debug日志输出
func SetDebugMode(is bool) {
	IsDebugMode = is
	Info("debug mode set to \n", is)
}

// SetEnv 设置环境
func SetEnv(env string) {
	Env = env
	Info("env set to", env)
}

// init log基础设置
func init() {
	file := "./" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[inputMethod]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

// Debug 打印调试日志
func Debug(format string, v ...interface{}) {
	if Env == EnvDevbox {
		fmt.Printf(time.Now().Format(DateTimeFmt3)+" "+format+"\n", v...)
	}
	log.Printf(format, v...)
}

// Err 错误日志
func Err(format string, v ...interface{}) {
	if Env == EnvDevbox {
		fmt.Printf(time.Now().Format(DateTimeFmt3) + " " + format)
		fmt.Printf("%+v\n", v...)
	}
	log.Print(format)
	log.Printf("ERROR:%+v\n", v...)

}

// Info 运行信息日志
func Info(format string, v ...interface{}) {
	if Env == EnvDevbox {
		fmt.Printf(time.Now().Format(DateTimeFmt3) + " " + format)
		fmt.Printf("%+v\n", v...)

	}
	log.Print(format)
	log.Printf("INFO:%+v\n", v...)
}

// Warning 警告日志
func Warning(format string, v ...interface{}) {
	if Env == EnvDevbox {
		fmt.Printf(time.Now().Format(DateTimeFmt3) + " " + format)
		fmt.Printf("%+v\n", v...)

	}
	log.Print(format)
	log.Printf("WARN:%+v\n", v...)

}
