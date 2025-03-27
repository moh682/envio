package domain

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type logger struct {
	z   *zap.Logger
	ctx context.Context
}

func NewLogger(ctx context.Context, z *zap.Logger) Logger {
	return &logger{
		z:   z,
		ctx: ctx,
	}
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.z.Info(msg, zap.Any("args", args))
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.z.Error(msg, zap.Any("args", args))
}
