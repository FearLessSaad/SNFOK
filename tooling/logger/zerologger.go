package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var (
	logger *zerolog.Logger
	once   sync.Once
)

const (
	INFO  = "INFO"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
	WARN  = "WARN"

	ERROR_MESSAGE = "err_msg"
)

// Field represents a key-value pair for structured logging.
type Field struct {
	Key   string
	Value interface{}
}

// Log writes a log message at the specified level with optional fields.
// If level is empty, it defaults to "info".
// Logs are written to both a file and stdout.
func Log(level, message string, fields ...Field) {
	// Initialize logger exactly once
	once.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		// Get configuration from environment variables
		logFile := os.Getenv("LOG_FILE")
		if logFile == "" {
			logFile = "app.log"
		}

		// Configure lumberjack for log rotation
		fileWriter := &lumberjack.Logger{
			Filename:   logFile, // Log file path
			MaxSize:    100,     // Rotate after 100 MB
			MaxBackups: 3,       // Keep 3 old log files
			MaxAge:     7,       // Retain logs for 7 days
			Compress:   true,    // Compress rotated files
		}

		// Configure console writer for human-readable output
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05", // Complete date-time (e.g., 2025-05-01 12:58:00)
			FormatLevel: func(i interface{}) string {
				switch i.(string) {
				case "debug":
					return "DEBUG"
				case "info":
					return "INFO"
				case "warn":
					return "WARN"
				case "error":
					return "ERROR"
				default:
					return "INFO"
				}
			},
			FormatFieldName: func(i interface{}) string {
				return fmt.Sprintf("%s=", i)
			},
			FormatFieldValue: func(i interface{}) string {
				return fmt.Sprintf("%v", i)
			},
		}

		// Create a multi-writer for file and stdout
		multiWriter := zerolog.MultiLevelWriter(
			fileWriter,
			consoleWriter,
		)

		// Create a single logger instance
		newLogger := zerolog.New(multiWriter).With().
			Timestamp().
			Logger()
		logger = &newLogger
	})

	// Default to info level if not specified
	if level == "" {
		level = INFO
	}

	// Log the message based on the level
	event := logger.Info() // Default to info
	switch level {
	case DEBUG:
		event = logger.Debug()
	case WARN:
		event = logger.Warn()
	case ERROR:
		event = logger.Error()
	}

	// Add optional fields
	for _, f := range fields {
		switch v := f.Value.(type) {
		case string:
			event.Str(f.Key, v)
		case int:
			event.Int(f.Key, v)
		case float64:
			event.Float64(f.Key, v)
		case bool:
			event.Bool(f.Key, v)
		default:
			event.Interface(f.Key, v)
		}
	}

	event.Msg(message)
}
