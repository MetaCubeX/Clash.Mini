package log

import (
	"bufio"
	"fmt"
	"runtime"
	"strings"
	"time"

	commonUtils "github.com/MetaCubeX/Clash.Mini/util/common"

	cLog "github.com/Dreamacro/clash/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var (
	RotateWriter       *rotatelogs.RotateLogs
	fileFolder         = "log"
	fileMainName       = "Clash.Mini"
	fileSuffix         = ".log"
	fileDatetimeFormat = `_%Y-%m-%d-%p`
	fileMaxAge         = 15 * 24 * time.Hour // 15 days
	fileRotationTime   = 12 * time.Hour      // half of the day

	logLevelFlagMap = map[string]string{
		cLog.DEBUG.String():   "DEBG",
		cLog.INFO.String():    "INFO",
		cLog.WARNING.String(): "WARN",
		cLog.ERROR.String():   "EROR",
	}
)

func init() {
	//if !app.Debug {
	//hook, err := logrusBugsnag.NewBugsnagHook()
	//if err != nil {
	//	panic(err)
	//}
	//cLog.Logger.Hooks.Add(hook)
	//}
	go runLog()
	//Debugln("test")
	//Infoln("test")
	//Warnln("test")
	//Errorln("test1")
	//Errorln("test: %s%s", "123", "456")
}

func runLog() {
	logPath := commonUtils.GetExecutablePath(fileFolder, fileMainName)

	var err error
	RotateWriter, err = rotatelogs.New(
		logPath+fileDatetimeFormat+fileSuffix,
		rotatelogs.WithLinkName(logPath+fileSuffix),
		rotatelogs.WithMaxAge(fileMaxAge),
		rotatelogs.WithRotationTime(fileRotationTime),
	)
	if err != nil {
		Infoln("[log] failed to create rotatelogs: %v", err)
		return
	}

	logWriter := bufio.NewWriter(RotateWriter)

	sub := cLog.Subscribe()
	defer cLog.UnSubscribe(sub)

	for elm := range sub {
		log := interface{}(elm).(*cLog.Event)
		if log.LogLevel < cLog.Level() {
			continue
		}

		_, err = logWriter.WriteString(fmt.Sprintf("%s [ %s ] %s\n",
			time.Now().Format("2006-01-02T15:04:05-07:00"),
			logLevelFlagMap[log.Type()], log.Payload))
		if err != nil {
			break
		}

		err := logWriter.Flush()
		if err != nil {
			break
		}
	}
}

// GetTraceFuncName 获取日志调用来源方法名
func GetTraceFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	fName := f.Name()
	return fName[strings.LastIndexByte(fName, '/')+1:]
}

// Debugln 调试日志行
func Debugln(format string, v ...interface{}) {
	//cLog.Debugln(fmt.Sprintf("[%s] %s", GetTraceFuncName(), format), v...)
	if len(v) == 0 {
		cLog.Debugln("%s", format)
	} else {
		cLog.Debugln(format, v...)
	}
}

// Infoln 信息日志行
func Infoln(format string, v ...interface{}) {
	//cLog.Infoln(fmt.Sprintf("[%s] %s", GetTraceFuncName(), format), v...)
	if len(v) == 0 {
		cLog.Infoln("%s", format)
	} else {
		cLog.Infoln(format, v...)
	}
}

// Warnln 警告日志行
func Warnln(format string, v ...interface{}) {
	//cLog.Warnln(fmt.Sprintf("[%s] %s", GetTraceFuncName(), format), v...)
	if len(v) == 0 {
		cLog.Warnln("%s", format)
	} else {
		cLog.Warnln(format, v...)
	}
}

// Errorln 错误日志行
func Errorln(format string, v ...interface{}) {
	//cLog.Errorln(fmt.Sprintf("[%s] %s", GetTraceFuncName(), format), v...)
	if len(v) == 0 {
		cLog.Errorln("%s", format)
	} else {
		cLog.Errorln(format, v...)
	}
}

// Panicln 异常日志行
func Panicln(format string, v ...interface{}) {
	var msg string
	if len(v) == 0 {
		msg = fmt.Sprintf("%s", format)
	} else {
		msg = fmt.Sprintf(format, v...)
	}
	//cLog.Errorln(fmt.Sprintf("[%s] %s", GetTraceFuncName(), msg))
	cLog.Errorln(msg)
	panic(msg)
}

// Fatalln 致命错误日志行
// 会直接退出，且不上报日志
//
// Deprecated: 需要上报日志时，使用 Panicln 替代
func Fatalln(format string, v ...interface{}) {
	cLog.Fatalln(format, v...)
}

func Level() cLog.LogLevel {
	return cLog.Level()
}

func SetLevel(newLevel cLog.LogLevel) {
	cLog.SetLevel(newLevel)
}
