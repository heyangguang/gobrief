package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Logger 日志对象
var Logger *zap.Logger

// InitLogger 初始化Logger
func InitLogger(filename string, maxSize, maxAge, maxBackup int, level string) (err error) {
	// 日志切割第三方包
	ws := getLogWriter(filename, maxSize, maxAge, maxBackup)
	encoderWriteFileJson := getEncoderWriteFileJson()
	encoderConsole := getEncoderConsole()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoderWriteFileJson, ws, l),
		zapcore.NewCore(encoderConsole, zapcore.AddSync(os.Stdout), l),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return
}

// 日志切割
func getLogWriter(filename string, maxSize, maxAge, maxBackup int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,
		// 单个文件最大尺寸，默认单位M
		MaxSize: maxSize,
		// 日志最大时间，单位天
		MaxAge: maxAge,
		// 备份日志的数量
		// 比如设置为10，那么无论超不超过最大天数也还是会删除
		MaxBackups: maxBackup,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 日志格式JSON写到文件里
func getEncoderWriteFileJson() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 人类识别时间
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 秒级间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 函数调用关系
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回JSON
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 日志格式打印控制台  可读性高
func getEncoderConsole() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 人类识别时间
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 秒级间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 函数调用关系
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回JSON
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}
