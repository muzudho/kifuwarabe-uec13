package coding_obj

// Log - ロガー
var Log Logger = *new(Logger)

// Gtp - 標準出力とログを一緒にしたもの
var Gtp StdoutLogWriter = *NewStdoutLogWriter(&Log)

// Console - エラー出力とログを一緒にしたもの
var Console StderrLogWriter = *NewStderrLogWriter(&Log)
