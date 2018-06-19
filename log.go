package log

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/xaionaro-go/errors"
)

var _logger *logger

type logger struct {
	*log.Logger
}

func init() {
	_logger = &logger{Logger: log.New(os.Stderr, fmt.Sprintf("[%v] ", os.Args[0]), 0)}
}

func (l *logger) getCurrentLine() string {
	stack := string(debug.Stack())
	stackLines := strings.Split(stack, "\n")
	startLineIdx0 := 0
	foundMe := false
	for idx, line := range stackLines {
		if strings.Index(line, "github.com/xaionaro-go/log") != -1 {
			startLineIdx0 = idx
			foundMe = true
			break
		}
	}
	if !foundMe {
		return ""
	}
	foundNotMe := false
	startLineIdx1 := 0
	for idx, line := range stackLines[startLineIdx0:] {
		if strings.Index(line, "github.com/xaionaro-go/log") == -1 {
			startLineIdx1 = idx
			foundNotMe = true
			break
		}
	}
	if !foundNotMe {
		return ""
	}
	theLine := stackLines[startLineIdx0 + startLineIdx1 + 1]
	theLine = strings.Split(strings.Trim(theLine, " \t"), " ")[0]
	thePath := strings.Split(theLine, "/")
	return thePath[len(thePath)-1]
}

func (l *logger) Fatal(v ...interface{}) {
	l.Logger.Fatal(append([]interface{}{"[fatal] "+l.getCurrentLine()+": "}, v...)...)
}

func (l *logger) Panic(v ...interface{}) {
	l.Logger.Panic(append([]interface{}{"[panic] "+l.getCurrentLine()+": "}, v...)...)
}

func (l *logger) Debug(v ...interface{}) {
	l.Logger.Print(append([]interface{}{"[debug] "+l.getCurrentLine()+": "}, v...)...)
}

func (l *logger) Debugln(v ...interface{}) {
	l.Logger.Println(append([]interface{}{"[debug] "+l.getCurrentLine()+":"}, v...)...)
}

func (l *logger) Debugf(fmt string, v ...interface{}) {
	l.Logger.Printf("[debug] "+l.getCurrentLine()+": "+fmt, v...)
}

func (l *logger) Warning(v ...interface{}) {
	l.Logger.Print(append([]interface{}{"[warning] "+l.getCurrentLine()+": "}, v...)...)
}

func (l *logger) Warningf(fmt string, v ...interface{}) {
	l.Logger.Printf("[warning] "+l.getCurrentLine()+": "+fmt, v...)
}

func (l *logger) Errorf(fmt string, v ...interface{}) {
	l.Logger.Printf("[error] "+l.getCurrentLine()+": "+fmt, v...)
}

func Fatal(v ...interface{}) {
	_logger.Fatal(v...)
}

func Panic(v ...interface{}) {
	_logger.Panic(v...)
}

func Debug(v ...interface{}) {
	_logger.Debug(v...)
}

func Debugln(v ...interface{}) {
	_logger.Debugln(v...)
}

func Debugf(fmt string, v ...interface{}) {
	_logger.Debugf(fmt, v...)
}

func Warning(v ...interface{}) {
	_logger.Warning(v...)
}

func Warningf(fmt string, v ...interface{}) {
	_logger.Warningf(fmt, v...)
}

func Errorf(fmt string, v ...interface{}) {
	_logger.Errorf(fmt, v...)
}

func (l *logger) anyWrapper(outFunc func(...interface{}), smartErrTmpl errors.SmartError, err error, args ...interface{}) error {
	if err == nil && len(args) == 0 {
		return nil
	}

	resultErr := smartErrTmpl.New(err, args...).SetCutOffFirstNLinesOfTraceback(13)
	outFunc(resultErr.ErrorShort())
	return resultErr
}

func (l *logger) WarningWrapper(smartErrTmpl errors.SmartError, err error, args ...interface{}) error {
	return l.anyWrapper(l.Warning, smartErrTmpl, err, args...)
}

func WarningWrapper(smartErrTmpl errors.SmartError, err error, args ...interface{}) error {
	return _logger.WarningWrapper(smartErrTmpl, err, args...)
}
