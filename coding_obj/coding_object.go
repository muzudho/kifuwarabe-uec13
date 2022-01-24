package coding_obj

// GtpLog - Gtp用ロガー
var GtpLog StdoutLogger = *new(StdoutLogger)

// Gtp - 標準出力とログを一緒にしたもの
var Gtp StdoutLogWriter = *NewStdoutLogWriter(&GtpLog)

// ConsoleLog - Console用ロガー
var ConsoleLog StderrLogger = *new(StderrLogger)

// Console - エラー出力とログを一緒にしたもの
var Console StderrLogWriter = *NewStderrLogWriter(&ConsoleLog)
