package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// slog.Groupでネストされたログを作成する
	logger.Info("Info Message", slog.Group("グループ1", "追加情報", "Infoレベルです"), slog.Group("グループ2", slog.String("追加情報", "Infoレベルだ")))
	// {"time":"2024-12-02T22:44:50.577887+09:00","level":"INFO","msg":"Info Message","グループ1":{"追加情報":"Infoレベルです"},"グループ2":{"追加情報":"Infoレベルだ"}}
}
