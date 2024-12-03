package main

import "github.com/CHlluanma/go-pkg/log"

func main() {
	// slog example
	{
		// slog.SetLogLoggerLevel(slog.LevelInfo)
		// slog.Debug("debug message")
		// slog.Info("info message")
		// slog.Warn("warn message")
		// slog.Error("error message")

		// l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// 	AddSource:   true,
		// 	Level:       slog.LevelInfo,
		// 	ReplaceAttr: nil,
		// }))
		// l.Debug("debug message")
		// l.Info("info message")
		// l.Warn("warn message")
		// l.Error("error message")

		// l.LogAttrs(context.Background(), slog.LevelInfo, "info message", slog.Any("key", "value"))

		// slog.SetDefault(l)
		// slog.Info("info message")
	}
	// 二次封装 SlogLogger
	{
		log.Info("info message")

		l := log.New(log.LevelInfo)
		l.Info("info message")
	}
}
