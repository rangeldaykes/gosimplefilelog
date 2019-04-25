package gosimplefilelog

import (
	"container/list"
)

// LogContext controls the indentation in logs
type LogContext struct {
	// Use to control FIFO - LIFO
	methodStack list.List
	// Cntrols the spaces ahead of the indentation
	logPrefix string
	// Tell me how many spaces
	contextCount int
}

// GetMessage return message with identation
func (lc *LogContext) GetMessage(message string) string {
	return lc.logPrefix + message
}

// Start - initiate a identation of log
func (lc *LogContext) Start(operation string) string {
	lc.methodStack.PushBack(operation)

	operation = "-> " + operation

	lc.contextCount++
	lc.logPrefix = "  " + lc.logPrefix

	return operation
}

// End - finalize a identation of log
func (lc *LogContext) End() string {
	var operation string

	if lc.contextCount == 0 {
		return ""
	}

	f := lc.methodStack.Back()
	lc.methodStack.Remove(f)

	operation = f.Value.(string)

	if lc.contextCount == 1 {
		operation = lc.GetMessage("<- " + operation)

		lc.contextCount--
		lc.logPrefix = lc.logPrefix[2:]

	} else {
		lc.contextCount--
		lc.logPrefix = lc.logPrefix[2:]

		operation = lc.GetMessage("<- " + operation)
	}

	return operation
}
