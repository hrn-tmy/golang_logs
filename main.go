package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With(slog.String("メッセージ", "共通ログ"))

	logger.Info("Info Message")
	// {"time":"2024-12-02T22:39:06.791661+09:00","level":"INFO","msg":"Info Message","メッセージ":"共通ログ"}
	logger.Error("Error Message")
	// {"time":"2024-12-02T22:39:06.792175+09:00","level":"ERROR","msg":"Error Message","メッセージ":"共通ログ"}
}
