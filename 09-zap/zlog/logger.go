package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger

// debug会输出到控制台和文件  release只输出到文件
var mode = "release"

func Ready(logMode string) {
	mode = logMode
	initLog()
	initSugar()
}
func initLog() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//日志级别
	//error级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	//info和debug级别,debug级别是最低的
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件WriteSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log",
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     180,   // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	})
	//error文件WriteSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     360,   // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	})

	//生成core
	//multiWriteSyncer := zapcore.NewMultiWriteSyncer(writerSyncer, zapcore.AddSync(os.Stdout)) //AddSync将io.Writer转换成WriteSyncer的类型
	//同时输出到控制台 和 指定的日志文件中,生产环境下把控制台去掉。

	var infoFileCore, errorFileCore zapcore.Core
	if mode == "debug" {
		infoFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
		errorFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)
	} else {
		infoFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), lowPriority)
		errorFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer), highPriority)
	}

	logger = zap.New(
		zapcore.NewTee(infoFileCore, errorFileCore),
		zap.AddCaller(),      //显示文件名和行号
		zap.AddCallerSkip(1), //打印调用行,因为重新封装的Info等方法，所以跳过一行
		//zap.AddStacktrace(zapcore.ErrorLevel), //大于等于这个级别的打印全部调用堆栈
	)
}

// SetMode 默认release只输出到文件，debug会输出到控制台和文件

func Debug(format string, fileds ...zapcore.Field) {
	logger.Debug(format, fileds...)
}

func Info(format string, fileds ...zapcore.Field) {
	logger.Info(format, fileds...)
}

func Warn(format string, fileds ...zapcore.Field) {
	logger.Warn(format, fileds...)
}

func Error(format string, fileds ...zapcore.Field) {
	logger.Error(format, fileds...)
}

func Panic(format string, fileds ...zapcore.Field) {
	logger.Panic(format, fileds...)
}
