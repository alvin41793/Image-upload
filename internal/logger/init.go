package logger

import (
	"fmt"
	"github.com/alvin41793/Image-upload/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Init() error {
	cfg := config.G().Log()
	if err := os.MkdirAll(cfg.Dir, 0755); err != nil {
		return err
	}

	filename := filepath.Join(cfg.Dir, fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02")))
	fileWriter, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	writeSyncer := zapcore.AddSync(fileWriter)

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	core := zapcore.NewCore(encoder, writeSyncer, parseLogLevel(cfg.Level))
	zl := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // zl 是 *zap.Logger 类型
	InitGlobal(&zapLogger{l: zl})

	go cleanOldLogs()

	return nil
}

func parseLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func cleanOldLogs() {
	ticker := time.NewTicker(24 * time.Hour)
	cfg := config.G().Log()
	for {
		<-ticker.C
		filepath.WalkDir(cfg.Dir, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if strings.HasPrefix(d.Name(), "app_") && strings.HasSuffix(d.Name(), ".log") {
				dateStr := strings.TrimSuffix(strings.TrimPrefix(d.Name(), "app_"), ".log")
				t, err := time.Parse("2006-01-02", dateStr)
				if err == nil && time.Since(t) > time.Duration(cfg.KeepDays)*24*time.Hour {
					_ = os.Remove(path)
				}
			}
			return nil
		})
	}
}
