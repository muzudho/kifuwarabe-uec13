package coding_obj

// StdoutLogger - 標準出力と一緒に使うロガー。
type StdoutLogger struct {
	printPath string
}

func (logger *StdoutLogger) SetPath(printPath string) {
	logger.printPath = printPath
}

// Print - ログファイルに書き込みます。 StdoutLogWriter から呼び出してください。
func (logger *StdoutLogger) Print(text string, args ...interface{}) {
	write(logger.printPath, text, args...)
}
