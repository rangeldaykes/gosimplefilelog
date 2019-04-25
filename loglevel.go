package gosimplefilelog

// LogLevel - indicates hierarchical level from the lowest to the highest
type LogLevel int

const (
	// LEVEL_LOG - information messages
	LEVEL_LOG LogLevel = 1 + iota
	// LEVEL_DEBUG - debugger messages
	LEVEL_DEBUG
	// LEVEL_ASSERT - check sanity messages
	LEVEL_ASSERT
	// LEVEL_ERROR - error messages
	LEVEL_ERROR
)
