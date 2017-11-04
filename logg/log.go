package logg

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/akawa1/gbox/errr"
	l4g "github.com/alecthomas/log4go"
)

const (
	DefaultFormat string = "[%D %T] [%L] (%S) %M"
)

var (
	logger l4g.Logger
)

func init() {
	logger = l4g.NewDefaultLogger(l4g.DEBUG)
}

func SetConsoleLogger(lvl byte, format ...string) {
	clw := l4g.NewConsoleLogWriter()
	if len(format) > 0 {
		clw.SetFormat(format[0])
	}
	logger.AddFilter("stdout", level(lvl), clw)
}

func SetSingleFileLogger(lvl byte, filename string, format ...string) {
	SetFileLogger(lvl, filename, false, false, 0, 0, format...)
}

func SetFileLogger(lvl byte, filename string, rotate, rttDaily bool, rttSize, rttLines int, format ...string) {
	flw := l4g.NewFileLogWriter(filename, rotate)
	if len(format) > 0 {
		flw.SetFormat(format[0])
	}
	flw.SetRotateDaily(rttDaily)
	flw.SetRotateSize(rttSize)
	flw.SetRotateLines(rttLines)
	logger.AddFilter("file", level(lvl), flw)
}

func D(tag string, arg0 interface{}, v ...interface{}) {
	log('D', tag, arg0, v...)
}

func I(tag string, arg0 interface{}, v ...interface{}) {
	log('I', tag, arg0, v...)
}

func W(tag string, arg0 interface{}, v ...interface{}) {
	log('W', tag, arg0, v...)
}

func E(tag string, arg0 interface{}, v ...interface{}) {
	log('E', tag, arg0, v...)
}

func C(tag string, arg0 interface{}, v ...interface{}) {
	log('C', tag, arg0, v...)
}

func log(lvl byte, tag string, arg0 interface{}, args ...interface{}) {
	if len(tag) == 0 {
		tag = "NOTAG"
	}

	var msg, pos string
	switch t := arg0.(type) {
	case *errr.Error:
		msg = t.Message
		pos = t.Position
	case error:
		msg = t.Error()
	case string:
		msg = fmt.Sprintf(t, args...)
	default:
		msg = fmt.Sprintf("%v", arg0) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...)
	}

	if len(pos) == 0 {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			pos = filepath.Base(file) + ":" + strconv.Itoa(line)
		} else {
			pos = "NOPOS"
		}
	}

	logger.Log(level(lvl), pos, tag+"\t"+msg)
}

func level(lvl byte) l4g.Level {
	switch lvl {
	case 'D':
		return l4g.DEBUG
	case 'T':
		return l4g.TRACE
	case 'I':
		return l4g.INFO
	case 'W':
		return l4g.WARNING
	case 'E':
		return l4g.ERROR
	case 'C':
		return l4g.CRITICAL
	default:
		return l4g.CRITICAL
	}
}
