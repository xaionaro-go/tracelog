package log

import (
	"fmt"
	"log"
	"os"

	"github.com/xaionaro-go/errors"
)

var _logger *logger

type logger struct {
	*log.Logger
}

func init() {
	_logger = &logger{Logger: log.New(os.Stderr, fmt.Sprintf("[%v] ", os.Args[0]), 0)}
}

func (l *logger) Fatal(v ...interface{}) {
	l.Logger.Fatal(append([]interface{}{"[fatal] "}, v...)...)
}

func (l *logger) Panic(v ...interface{}) {
	l.Logger.Panic(append([]interface{}{"[panic] "}, v...)...)
}

func (l *logger) Debug(v ...interface{}) {
	l.Logger.Print(append([]interface{}{"[debug] "}, v...)...)
}

func (l *logger) Debugln(v ...interface{}) {
	l.Logger.Println(v...)
}

func (l *logger) Debugf(fmt string, v ...interface{}) {
	l.Logger.Printf(fmt, v...)
}

func (l *logger) Warning(v ...interface{}) {
	l.Logger.Print(v...)
}

func (l *logger) Warningf(fmt string, v ...interface{}) {
	l.Logger.Printf(fmt, v...)
}

func (l *logger) Errorf(fmt string, v ...interface{}) {
	l.Logger.Printf(fmt, v...)
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
