// 標準出力とロガーを一緒にしただけのもの
package coding_obj

import (
	"fmt"
)

// StdoutLogWriter - 標準出力とロガーを一緒にしただけです
type StdoutLogWriter struct {
	logger *Logger
}

// NewStdoutLogWriter - オブジェクト作成
func NewStdoutLogWriter(logger *Logger) *StdoutLogWriter {
	writer := new(StdoutLogWriter)
	writer.logger = logger
	return writer
}

// Print - 必ず出力します。
func (writer StdoutLogWriter) Print(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Print(text, args...) // ログ
}
