package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var sugarLogger *zap.SugaredLogger

func initSugar() {
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
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
		Filename:   "./logf/info.log",
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     180,   // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	})
	//error文件WriteSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logf/error.log",
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

	log := zap.New(
		zapcore.NewTee(infoFileCore, errorFileCore),
		zap.AddCaller(),      //显示文件名和行号
		zap.AddCallerSkip(1), //打印调用行,因为重新封装的Info等方法，所以跳过一行
		//zap.AddStacktrace(zapcore.ErrorLevel), //大于等于这个级别的打印全部调用堆栈
	)
	sugarLogger = log.Sugar()
}

func Debugf(format string, v ...interface{}) {
	sugarLogger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	sugarLogger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	sugarLogger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	sugarLogger.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	sugarLogger.Panicf(format, v...)
}
