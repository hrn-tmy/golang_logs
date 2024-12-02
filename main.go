package main

import (
	"log/slog"
	"os"
)

func main() {
	// 出力のカスタマイズを行う
	opt := slog.HandlerOptions{
		// ログの出力場所を詳述する
		AddSource: true,
		// 出力するログレベルを設定
		Level: slog.LevelWarn,
		// 追加情報の出力に関する設定
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Value.String() == "Warn Message" {
				a.Value = slog.StringValue("WARN MESSAGE")
			}
			return a
		},
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	// Warn以上を出力するように設定したため出力されない
	logger.Info("Info Message")
	logger.Warn("Warn Message")
	// {"time":"2024-12-02T22:29:02.602679+09:00","level":"WARN","source":{"function":"main.main","file":"/Users/hoge/fuga/golang_logs/main.go","line":21},"msg":"WARN MESSAGE"}
	logger.Error("Error Message")
	// {"time":"2024-12-02T22:29:02.603285+09:00","level":"ERROR","source":{"function":"main.main","file":"/Users/hoge/fuga/golang_logs/main.go","line":22},"msg":"Error Message"}
}
