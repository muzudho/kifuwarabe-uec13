package coding_obj

// GtpLog - Gtp用ロガー
var GtpLog Logger = *new(Logger)

// Gtp - 標準出力とログを一緒にしたもの
var Gtp StdoutLogWriter = *NewStdoutLogWriter(&GtpLog)

// ConsoleLog - Console用ロガー
var ConsoleLog Logger = *new(Logger)

// Console - エラー出力とログを一緒にしたもの
var Console StderrLogWriter = *NewStderrLogWriter(&ConsoleLog)
