package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for
	// debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a loggin level that contains information deemed
	// valuable
	LevelInfo
	// LevelError represents the highest logging level, only to be used to
	// trace errors.
	LevelError
)

func (lvl Level) String() string {
	switch lvl {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelError:
		return "[ERROR]"
	default:
		return ""
	}
}
