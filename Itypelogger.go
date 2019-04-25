package gosimplefilelog

// ITypeLogger - indicates what it takes to be a logger
type ITypeLogger interface {
	Log(level LogLevel, message string)
}
