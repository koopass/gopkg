package zlog

import (
	"context"

	"go.uber.org/zap"
)

var defaultLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defaultLogger = logger.Sugar()
}

func Init(cfg *zap.Config) error {
	logger, err := cfg.Build()
	if err != nil {
		return err
	}
	defaultLogger = logger.Sugar()
	return nil
}

func CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func CtxInfof(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

func CtxWarnf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

func CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func CtxFatalf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

func CtxPanicf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Panicf(template, args...)
}

func CtxDebugw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Debugw(template, args...)
}

func CtxInfow(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Infow(template, args...)
}

func CtxWarnw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Warnw(template, args...)
}

func CtxErrorw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Errorw(template, args...)
}

func CtxFatalw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Fatalw(template, args...)
}

func CtxPanicw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Panicw(template, args...)
}

func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	defaultLogger.Panicf(template, args...)
}

func Debugw(template string, args ...interface{}) {
	defaultLogger.Debugw(template, args...)
}

func Infow(template string, args ...interface{}) {
	defaultLogger.Infow(template, args...)
}

func Warnw(template string, args ...interface{}) {
	defaultLogger.Warnw(template, args...)
}

func Errorw(template string, args ...interface{}) {
	defaultLogger.Errorw(template, args...)
}

func Fatalw(template string, args ...interface{}) {
	defaultLogger.Fatalw(template, args...)
}

func Panicw(template string, args ...interface{}) {
	defaultLogger.Panicw(template, args...)
}
