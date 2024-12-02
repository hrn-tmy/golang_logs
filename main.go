package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Info Message", slog.String("追加メッセージ", "Infoレベルのメッセージです"))
	logger.Info("Info Message", "追加メッセージ", "Infoレベルのメッセージです")
	// {"time":"2024-12-02T22:17:38.192361+09:00","level":"INFO","msg":"Info Message","追加メッセージ":"Infoレベルのメッセージです"}

	logger.Info("Info Message", "追加メッセージ")
	// {"time":"2024-12-02T22:18:28.502302+09:00","level":"INFO","msg":"Info Message","!BADKEY":"追加メッセージ"}
}
