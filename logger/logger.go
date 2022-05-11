package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	infoLog *log.Logger
	warnLog *log.Logger
	errLog  *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, color.CyanString("INFO "), log.Ltime|log.Ldate)
	warnLog = log.New(os.Stdout, color.YellowString("WARN "), log.Ltime|log.Ldate)
	errLog = log.New(os.Stdout, color.HiRedString("ERR "), log.Ltime|log.Ldate)
}

func Info(ctx ...interface{}) {
	infoLog.Println(ctx...)
}

func Warn(ctx ...interface{}) {
	warnLog.Println(ctx...)
}

func Err(ctx ...interface{}) {
	errLog.Println(ctx...)
}
