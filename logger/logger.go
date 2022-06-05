package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errLog   *log.Logger
)

func init() {
	debugLog = log.New(
		os.Stdout,
		color.YellowString("DEBUG "),
		log.Ltime|log.Ldate)

	infoLog = log.New(
		os.Stdout,
		color.New(color.FgHiGreen).Sprintf("INFO "),
		log.Ltime|log.Ldate)

	warnLog = log.New(
		os.Stdout,
		color.New(color.Bold, color.FgHiYellow).Sprintf("WARN "),
		log.Ltime|log.Ldate)

	errLog = log.New(
		os.Stdout,
		color.New(color.Bold, color.FgRed).Sprintf("ERROR "),
		log.Ltime|log.Ldate)
}

func Debug(ctx ...interface{}) {
	debugLog.Println(ctx...)
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
