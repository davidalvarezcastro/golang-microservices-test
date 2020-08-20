package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log is the log variable
	Log *zap.Logger
)

func init() {
	var err error
	logConfig := zap.Config{
		OutputPaths: []string{"stdout", "/tmp/logs"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	Log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

// Field returns a Field from a string key
func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// Debug logs an debug message
func Debug(msg string, tags ...zap.Field) {
	defer Log.Sync()
	Log.Debug(msg, tags...)
}

// Info logs an info message
func Info(msg string, tags ...zap.Field) {
	defer Log.Sync()
	Log.Info(msg, tags...)
}

// Error logs an error message
func Error(msg string, err error, tags ...zap.Field) {
	defer Log.Sync()
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.Error(msg, tags...)
}
