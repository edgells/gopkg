package logger

import "context"

type Level int8

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = iota - 1
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
	// LevelFatal is logger fatal level
	LevelFatal
)

type Logger interface {
	Log(level Level, kv ...interface{})
}

type logger struct {
	logger    Logger
	prefix    []interface{}
	hasValuer bool
	ctx       context.Context
}
