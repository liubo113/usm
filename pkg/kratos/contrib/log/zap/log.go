// zap 实现 kratos 的 log.Logger 接口，支持日志轮转、syslog、dev 模式
package zap

import (
	"fmt"
	"log/syslog"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ log.Logger = (*Logger)(nil)

const timeLayout = "2006-01-02 15:04:05.999999"

type Logger struct {
	file  *lumberjack.Logger
	sugar *zap.SugaredLogger
}

type Config struct {
	Dev               bool          // 是否为开发模式，开发模式将日志输出到 stdout 中
	LogFilePath       string        // 日志文件存储路径
	LogFileMaxSize    int           // 单个日志文件最大大小，单位为 MB, 默认: 10MB
	LogFileMaxBackups int           // 日志轮转支持的最多归档数量，默认为：10 个
	Prefix            string        // 日志前缀
	AddStacktrace     bool          // 是否打印 stacktrace
	StacktraceLevel   zapcore.Level // stacktrace 级别
}

func NewLogger(conf *Config) (*Logger, error) {
	if conf == nil {
		conf = &Config{}
	}
	if conf.LogFilePath == "" {
		conf.LogFilePath = "./default.log"
	}
	if conf.LogFileMaxSize <= 0 {
		conf.LogFileMaxSize = 10
	}
	if conf.LogFileMaxBackups <= 0 {
		conf.LogFileMaxBackups = 1
	}
	l := &Logger{}
	if err := l.init(conf); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.sugar.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	switch level {
	case log.LevelDebug:
		l.sugar.Debug(keyvals...)
	case log.LevelInfo:
		l.sugar.Info(keyvals...)
	case log.LevelWarn:
		l.sugar.Warn(keyvals...)
	case log.LevelError:
		l.sugar.Error(keyvals...)
	case log.LevelFatal:
		l.sugar.Fatal(keyvals...)
	}
	return nil
}

func (l *Logger) init(conf *Config) error {
	var (
		zapConf zap.Config
		wss     []zapcore.WriteSyncer
	)
	if conf.Dev {
		zapConf = zap.NewDevelopmentConfig()
		ws, _, err := zap.Open(zapConf.OutputPaths...)
		if err != nil {
			return err
		}
		wss = append(wss, ws)
	} else {
		zapConf = zap.NewProductionConfig()
		l.file = &lumberjack.Logger{
			Filename:   conf.LogFilePath,
			MaxSize:    conf.LogFileMaxSize,
			MaxBackups: conf.LogFileMaxBackups,
		}
		ws := zapcore.AddSync(l.file)
		wss = append(wss, ws)
		syslogWriter, err := syslog.New(syslog.LOG_ERR, "")
		if err == nil {
			wss = append(wss, zapcore.AddSync(syslogWriter))
		}
	}
	zapConf.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeLayout))
	}
	zapConf.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var sb strings.Builder
		if conf.Prefix != "" {
			sb.WriteString(fmt.Sprintf("[%s] ", conf.Prefix))
		}
		sb.WriteString(fmt.Sprintf("[%s]", level.CapitalString()))
		enc.AppendString(sb.String())
	}
	zapConf.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%s:", caller.TrimmedPath()))
	}
	enc := zapcore.NewConsoleEncoder(zapConf.EncoderConfig)
	var zapCore []zapcore.Core
	for _, ws := range wss {
		zapCore = append(zapCore, zapcore.NewCore(&encoderWrapper{Encoder: enc}, ws, zapConf.Level))
	}
	base := zap.New(zapcore.NewTee(zapCore...))
	if conf.AddStacktrace {
		base = base.WithOptions(zap.AddStacktrace(conf.StacktraceLevel))
	}
	base = base.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	l.sugar = base.Sugar()
	return nil
}

type encoderWrapper struct {
	zapcore.Encoder
}

func (ew *encoderWrapper) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := ew.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	bs := buf.String()
	bs = strings.ReplaceAll(bs, string('\t'), " ")
	buf.Reset()
	buf.Write([]byte(bs))
	return buf, nil
}

func (l *Logger) Sync() error {
	err := l.sugar.Sync()
	if l.file != nil {
		err = l.file.Close()
	}
	return err
}
