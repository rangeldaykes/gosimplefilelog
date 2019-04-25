package gosimplefilelog

import (
	"fmt"
	"runtime/debug"
)

// Logger is expose methods to use log
type Logger struct {
	// Instance of type of output
	instance ITypeLogger
	// Hierarchical level of logs
	loglevel LogLevel
	// Used to create indentation in the log
	context LogContext
}

// NewLogger Creates a new logger.
// instantiate your own:
//
//	log := mylogger1.NewLogger(
//		mylogger1.NewFileLogger(
//			"Logtest",
//			"/var/logs/"),
//		mylogger1.LEVEL_LOG)
//
func NewLogger(typelogger ITypeLogger, Level LogLevel) *Logger {
	logger := &Logger{instance: typelogger, loglevel: Level, context: LogContext{}}
	return logger
}

// Log used to create a simple info message
func (l *Logger) Log(message string) {
	l.logBase(LEVEL_LOG, message)
}

// Debug used to cretate debug message to help implementation
func (l *Logger) Debug(message string) {
	l.logBase(LEVEL_DEBUG, message)
}

// Assert used to create sanity test messages
func (l *Logger) Assert(condition bool, message string) {
	if !condition {
		msg := fmt.Sprintf("%s %s", "Assert: ", message)
		l.logBase(LEVEL_ASSERT, msg)
	}
}

// Error used to create errors messages
func (l *Logger) Error(err error, message string) {

	msg := fmt.Sprintf("%s - %s \n %s \n %s", "Erro: ", message, err.Error(), debug.Stack())

	l.logBase(LEVEL_ERROR, msg)
}

// ContextStart used to initialize indentation log
func (l *Logger) ContextStart(message string) {
	message = l.context.Start(message)
	l.logBase(LEVEL_LOG, message)
}

// ContextEnd used to finalize the log indentation
func (l *Logger) ContextEnd() {
	message := l.context.End()
	l.logBase(LEVEL_LOG, message)
}

// logBase call writer log
func (l *Logger) logBase(level LogLevel, message string) {
	if l.loglevel > level {
		return
	}

	if l.instance != nil {
		if message != "" {
			message = ": " + l.context.GetMessage(message)
		}

		l.instance.Log(level, message)
	}
}
