package logger

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var loggerObj *kitexzap.Logger

type Logger struct {
	*kitexzap.Logger
}

func init() {
	loggerObj = DefaultLogger()
	klog.SetLogger(loggerObj)
}

func (l *Logger) Printf(template string, args ...interface{}) {
	l.Infof(template, args...)
}

func GetLogger() *Logger {
	return &Logger{loggerObj}
}

const (
	DefaultSkip = 4 // 默认跳过的栈帧数
)

type Config struct {
	Enc zapcore.Encoder
	Ws  zapcore.WriteSyncer
	lvl zapcore.Level
}

func InitLoggerWithLevel(lvl zapcore.Level) {
	klog.SetLogger(NewLogger(lvl, Config{}))
}

func InitLoggerWithConfig(lvl zapcore.Level, cfg Config, options ...zap.Option) {
	klog.SetLogger(NewLogger(lvl, cfg, options...))
}

func NewLogger(lvl zapcore.Level, cfg Config, options ...zap.Option) *kitexzap.Logger {
	if cfg.Enc == nil {
		cfg.Enc = defaultEnc()
	}
	if cfg.Ws == nil {
		cfg.Ws = defaultWs()
	}
	cfg.lvl = lvl

	var ops []kitexzap.Option
	ops = append(ops, kitexzap.WithZapOptions(defaultOptions()...))
	ops = append(ops, kitexzap.WithCoreEnc(cfg.Enc))
	ops = append(ops, kitexzap.WithCoreWs(cfg.Ws))
	ops = append(ops, kitexzap.WithCoreLevel(zap.NewAtomicLevelAt(cfg.lvl)))
	ops = append(ops, kitexzap.WithZapOptions(options...))
	return kitexzap.NewLogger(ops...)
}

func DefaultLogger(options ...zap.Option) *kitexzap.Logger {
	var ops []kitexzap.Option
	ops = append(ops, kitexzap.WithZapOptions(defaultOptions()...))
	ops = append(ops, kitexzap.WithCoreEnc(defaultEnc()))
	ops = append(ops, kitexzap.WithCoreWs(defaultWs()))
	ops = append(ops, kitexzap.WithCoreLevel(zap.NewAtomicLevelAt(defaultLvl())))
	ops = append(ops, kitexzap.WithZapOptions(options...))
	return kitexzap.NewLogger(ops...)
}

func defaultEnc() zapcore.Encoder {
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志等级大写
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(cfg)
}

func defaultWs() zapcore.WriteSyncer {
	return os.Stdout
}

func defaultLvl() zapcore.Level {
	return zapcore.DebugLevel
}

func defaultOptions() []zap.Option {
	return []zap.Option{
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(DefaultSkip),
	}
}

func Debug(args ...interface{}) {
	klog.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	klog.Debugf(template, args...)
}

func Info(args ...interface{}) {
	klog.Info(args...)
}

func Infof(template string, args ...interface{}) {
	klog.Infof(template, args...)
}

func Warn(args ...interface{}) {
	klog.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	klog.Warnf(template, args...)
}

func Error(args ...interface{}) {
	klog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	klog.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	klog.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	klog.Fatalf(template, args...)
}

// LErrorf Equals Errorf less one stack
func LErrorf(template string, args ...interface{}) {
	loggerObj.Errorf(template, args...)
}
