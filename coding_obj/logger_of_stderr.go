package coding_obj

// StderrLogger - エラー出力と一緒に使うロガー。
type StderrLogger struct {
	tracePath  string
	debugPath  string
	infoPath   string
	noticePath string
	warnPath   string
	errorPath  string
	fatalPath  string
	printPath  string
}

func (logger *StderrLogger) SetPath(
	tracePath string,
	debugPath string,
	infoPath string,
	noticePath string,
	warnPath string,
	errorPath string,
	fatalPath string,
	printPath string) {

	logger.tracePath = tracePath
	logger.debugPath = debugPath
	logger.infoPath = infoPath
	logger.noticePath = noticePath
	logger.warnPath = warnPath
	logger.errorPath = errorPath
	logger.fatalPath = fatalPath
	logger.printPath = printPath
}

// Trace - ログファイルに書き込みます。
func (logger *StderrLogger) Trace(text string, args ...interface{}) {
	write(logger.tracePath, text, args...)
}

// Debug - ログファイルに書き込みます。
func (logger *StderrLogger) Debug(text string, args ...interface{}) {
	write(logger.debugPath, text, args...)
}

// Info - ログファイルに書き込みます。
func (logger *StderrLogger) Info(text string, args ...interface{}) {
	write(logger.infoPath, text, args...)
}

// Notice - ログファイルに書き込みます。
func (logger *StderrLogger) Notice(text string, args ...interface{}) {
	write(logger.noticePath, text, args...)
}

// Warn - ログファイルに書き込みます。
func (logger *StderrLogger) Warn(text string, args ...interface{}) {
	write(logger.warnPath, text, args...)
}

// Error - ログファイルに書き込みます。
func (logger *StderrLogger) Error(text string, args ...interface{}) {
	write(logger.errorPath, text, args...)
}

// Fatal - ログファイルに書き込みます。
func (logger *StderrLogger) Fatal(text string, args ...interface{}) {
	write(logger.fatalPath, text, args...)
}

// Print - ログファイルに書き込みます。 StdoutLogWriter から呼び出してください。
func (logger *StderrLogger) Print(text string, args ...interface{}) {
	write(logger.printPath, text, args...)
}
