package tlog

import (
	"fmt"
	"log/slog"
	"os"
)

func Infof(message string, v ...any) {
	slog.Info(fmt.Sprintf(message, v...))
}

func Warnf(message string, v ...any) {
	slog.Warn(fmt.Sprintf(message, v...))
}

func Fatalf(message string, v ...any) {
	slog.Error(fmt.Sprintf(message, v...))
	os.Exit(1)
}

func Panicf(message string, v ...any) {
	e := fmt.Sprintf(message, v...)
	slog.Error(e)
	panic(e)
}

func Debugf(message string, v ...any) {
	slog.Debug(fmt.Sprintf(message, v...))
}

func Info(v ...any) {
	slog.Info(string(fmt.Appendln(nil, v...)))
}

func Warn(v ...any) {
	slog.Warn(string(fmt.Appendln(nil, v...)))
}

func Fatal(v ...any) {
	slog.Error(string(fmt.Appendln(nil, v...)))
	os.Exit(1)
}

func Panic(v ...any) {
	e := string(fmt.Appendln(nil, v...))
	slog.Error(e)
	panic(e)
}

func Debug(v ...any) {
	slog.Debug(string(fmt.Appendln(nil, v...)))
}

func SetLevelDebug() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func SetLevelInfo() {
	slog.SetLogLoggerLevel(slog.LevelInfo)
}

func SetLevelWarn() {
	slog.SetLogLoggerLevel(slog.LevelWarn)
}

func SetLevelError() {
	slog.SetLogLoggerLevel(slog.LevelError)
}
