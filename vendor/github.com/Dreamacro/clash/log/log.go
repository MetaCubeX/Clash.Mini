package log

import (
	"fmt"
	"io"
	"os"

	"github.com/Dreamacro/clash/common/observable"

	log "github.com/sirupsen/logrus"
)

var (
	logCh  = make(chan interface{})
	source = observable.NewObservable(logCh)
	level  = INFO
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	//InitExLog(log.DebugLevel, false, "./log.txt")
}

func InitExLog(logLevel log.Level, useJsonFormat bool, logFileName string) {
	log.SetLevel(logLevel)
	if useJsonFormat {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
	log.SetOutput(os.Stdout)
	
	writers := []io.Writer{ os.Stdout }
	var file *os.File
	var err error
	if len(logFileName) > 0 {
		log.Infof("log to file: %s", logFileName)
		file, err = os.OpenFile(logFileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
		if err != nil {
			log.Errorf("cannot log to file: %s", logFileName)
		} else {
			writers = append(writers, file)
		}
	}
	if len(writers) > 1 {
		multiWriter := io.MultiWriter(writers...)
		if err == nil {
			log.SetOutput(multiWriter)
		} else {
			log.Errorf("cannot log to multi writers: %v", writers)
		}
	}
}

type Event struct {
	LogLevel LogLevel
	Payload  string
}

func (e *Event) Type() string {
	return e.LogLevel.String()
}

func Infoln(format string, v ...interface{}) {
	event := newLog(INFO, format, v...)
	logCh <- event
	print(event)
}

func Warnln(format string, v ...interface{}) {
	event := newLog(WARNING, format, v...)
	logCh <- event
	print(event)
}

func Errorln(format string, v ...interface{}) {
	event := newLog(ERROR, format, v...)
	logCh <- event
	print(event)
}

func Debugln(format string, v ...interface{}) {
	event := newLog(DEBUG, format, v...)
	logCh <- event
	print(event)
}

func Fatalln(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func Subscribe() observable.Subscription {
	sub, _ := source.Subscribe()
	return sub
}

func UnSubscribe(sub observable.Subscription) {
	source.UnSubscribe(sub)
}

func Level() LogLevel {
	return level
}

func SetLevel(newLevel LogLevel) {
	level = newLevel
}

func print(data *Event) {
	if data.LogLevel < level {
		return
	}

	switch data.LogLevel {
	case INFO:
		log.Infoln(data.Payload)
	case WARNING:
		log.Warnln(data.Payload)
	case ERROR:
		log.Errorln(data.Payload)
	case DEBUG:
		log.Debugln(data.Payload)
	}
}

func newLog(logLevel LogLevel, format string, v ...interface{}) *Event {
	return &Event{
		LogLevel: logLevel,
		Payload:  fmt.Sprintf(format, v...),
	}
}
