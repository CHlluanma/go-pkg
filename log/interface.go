package log

import "context"

type InfoLogger interface {
	Info(msg string, args ...any)
	Infof(format string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)

	Enabled() bool
}

type Logger interface {
	InfoLogger
	Debug(ctx context.Context, msg string, args ...any)
	Debugf(ctx context.Context, format string, args ...any)
	Debugw(ctx context.Context, msg string, args ...any)

	Trace(ctx context.Context, msg string, args ...any)
	Tracef(ctx context.Context, format string, args ...any)
	Tracew(ctx context.Context, msg string, args ...any)

	Warn(ctx context.Context, msg string, args ...any)
	Warnf(ctx context.Context, format string, args ...any)
	Warnw(ctx context.Context, msg string, args ...any)

	Error(ctx context.Context, msg string, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
	Errorw(ctx context.Context, msg string, args ...any)

	// V 返回一个具有指定日志级别的信息记录器。
	// 这个方法允许动态调整日志记录的详细程度，以根据需要查看更具体或更高级别的日志信息。
	// 返回的 InfoLogger 是一个具有指定日志级别并可用于记录信息的接口。
	V(level Level) InfoLogger

	With(args ...any) Logger
	WithName(name string) Logger
}
