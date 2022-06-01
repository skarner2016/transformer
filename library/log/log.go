package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"transformer/library/config"
	"path"
	"runtime"
	"time"
)

var instant *zap.SugaredLogger

func NewLogger() *zap.SugaredLogger {
	if instant != nil {
		return instant
	}

	logConf := new(Conf)
	if err := config.VipConfig.UnmarshalKey(logKeyPrefix, &logConf); err != nil {
		panic(fmt.Sprintf("parse log config error:%v", err.Error()))
	}

	logFile := getCurrencyDir() + logConf.Path + "/" + getFile()
	writeSyncer := getWriter(logConf, logFile)
	encoder := getEncoder()
	level := getLogLevel(logConf.Level)

	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller())

	instant := logger.Sugar()

	return instant
}

func getWriter(c *Conf, logFile string) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackup,
		LocalTime:  c.LocalTime,
		Compress:   c.Compress,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getCurrencyDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	return path.Dir(filename) + "/../../"
}

func getFile() string {
	return time.Now().Format("2006-01-02") + ".log"
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogLevel(level string) zapcore.LevelEnabler {

	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	case "info":
	default:
		return zapcore.InfoLevel
	}

	return zapcore.InfoLevel
}
