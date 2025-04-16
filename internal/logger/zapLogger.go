package logger

import "go.uber.org/zap"

type zapLogger struct {
	l *zap.Logger
}

func (z *zapLogger) Debug(msg string, fields ...zap.Field) {
	z.l.Debug(msg, fields...)
}

func (z *zapLogger) Info(msg string, fields ...zap.Field) {
	z.l.Info(msg, fields...)
}

func (z *zapLogger) Warn(msg string, fields ...zap.Field) {
	z.l.Warn(msg, fields...)
}

func (z *zapLogger) Error(msg string, fields ...zap.Field) {
	z.l.Error(msg, fields...)
}

func (z *zapLogger) With(fields ...zap.Field) Logger {
	return &zapLogger{l: z.l.With(fields...)}
}
