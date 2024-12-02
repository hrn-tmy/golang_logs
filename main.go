package main

import (
	"log/slog"
	"os"
)

func main() {
	// デフォルトではInfo以上のログが出力されるので、Debugレベルは出力されない
	slog.Debug("Plain Debug")
	slog.Info("Plain Info")
	// 2024/12/02 21:55:21 INFO Plain Info
	slog.Warn("Plain Warn")
	// 2024/12/02 21:55:21 WARN Plain Warn
	slog.Error("Plain Error")
	// 2024/12/02 21:55:21 Error Plain Warn

	// カスタムロガー(JSON形式)の作成
	jsonSlog := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// デフォルトではInfo以上のログが出力されるので、Debugレベルは出力されない
	jsonSlog.Debug("JSON Debug Message")
	jsonSlog.Info("JSON Info Message")
	// {"time":"2024-12-02T21:55:21.508998+09:00","level":"INFO","msg":"Custom Info Message"}
	jsonSlog.Warn("JSON Warn Message")
	// {"time":"2024-12-02T21:55:21.509004+09:00","level":"WARN","msg":"Custom Warn Message"}
	jsonSlog.Error("JSON Error Message")
	// {"time":"2024-12-02T21:55:21.509008+09:00","level":"ERROR","msg":"Custom Error Message"}

	// カスタムロガー(テキスト形式)の作成
	textSlog := slog.New(slog.NewTextHandler(os.Stdout,nil))
	// デフォルトではInfo以上のログが出力されるので、Debugレベルは出力されない
	textSlog.Debug("JSON Debug Message")
	textSlog.Info("Text Info Message")
	// time=2024-12-02T21:59:05.051+09:00 level=INFO msg="Text Info Message"
	textSlog.Warn("Text Warn Message")
   	// time=2024-12-02T21:59:05.051+09:00 level=WARN msg="Text Warn Message"
	textSlog.Error("Text Error Message")
	// time=2024-12-02T21:59:05.051+09:00 level=ERROR msg="Text Error Message"
}