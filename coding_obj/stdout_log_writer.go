// 標準出力とロガーを一緒にしただけのもの
package coding_obj

import (
	"fmt"
)

// StdoutLogWriter - 標準出力とロガーを一緒にしただけです
type StdoutLogWriter struct {
	logger *StdoutLogger
}

// NewStdoutLogWriter - オブジェクト作成
func NewStdoutLogWriter(logger *StdoutLogger) *StdoutLogWriter {
	writer := new(StdoutLogWriter)
	writer.logger = logger
	return writer
}

// Print - 必ず出力します。
func (writer *StdoutLogWriter) Print(text string, args ...interface{}) {
	fmt.Printf(text, args...)          // 標準出力
	writer.logger.Print(text, args...) // ログ
}

// Log - ログだけ
func (writer *StdoutLogWriter) Log(text string, args ...interface{}) {
	writer.logger.Print(text, args...) // ログ
}
