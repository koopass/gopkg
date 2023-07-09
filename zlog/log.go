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

func Debugf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

func Panicf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Panicf(template, args...)
}

func Debugw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Debugw(template, args...)
}

func Infow(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Infow(template, args...)
}

func Warnw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Warnw(template, args...)
}

func Errorw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Errorw(template, args...)
}

func Fatalw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Fatalw(template, args...)
}

func Panicw(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Panicw(template, args...)
}
