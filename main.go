package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout,nil))
	// 上記ログをデフォルトとして指定することができる
	slog.SetDefault(logger)
	slog.Info("Info Message")
	// {"time":"2024-12-02T22:06:41.895091+09:00","level":"INFO","msg":"Info Message"}
	slog.Warn("Warn Message")
	// {"time":"2024-12-02T22:06:41.895399+09:00","level":"WARN","msg":"Warn Message"}
	slog.Error("Error Message")
	// {"time":"2024-12-02T22:06:41.895403+09:00","level":"ERROR","msg":"Error Message"}
}