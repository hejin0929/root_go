package my_log

import (
	"fmt"
	"log"
	"os"
)

func WriteLog() *log.Logger {
	logFile, err := os.OpenFile(".\\src\\static\\log\\reop_test.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logFile.Close()

	logger := log.New(logFile, "Rope", log.Ldate|log.Ltime|log.Llongfile)

	return logger
}

// ErrorLog
// 重大bug记录日志器
type ErrorLog struct {
	LogFile *os.File `json:"log_file"`
}

//var ErrLog ErrorLog

//func init() {
//}

func (_this ErrorLog) GetLogger() *log.Logger {

	fmt.Println(_this.LogFile)

	logger := log.New(_this.LogFile, "LogError --> ", log.Ldate|log.Ltime|log.Llongfile)

	return logger
}

func (_this ErrorLog) CloseFileLog() {
	defer _this.LogFile.Close()
}

func GetLog() ErrorLog {

	ErrLog := ErrorLog{}

	var err error

	ErrLog.LogFile, err = os.OpenFile(".\\src\\static\\log\\test_err.log", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	return ErrLog
}

//func WriteLogErrs(str string) *log.Logger {
//
//	logFile, err := os.OpenFile(".\\src\\static\\log\\test_err.log",os.O_RDWR | os.O_CREATE,0666)
//
//	if err != nil {
//		fmt.Printf("%s\r\n", err.Error())
//		os.Exit(-1)
//	}
//
//	logger := log.New(logFile,"LogErr  ", log.Ldate | log.Ltime | log.Llongfile)
//
//	defer logFile.Close()
//
//	return logger
//}
