package server

// Command type
type Command string

// command constants
const (
	CREATE = Command("CREATE")
	DROP   = Command("DROP")
	LIST   = Command("LIST")
	INFO   = Command("INFO")
	SET    = Command("SET")
	CHECK  = Command("CHECK")

	EndOfLine = '\n'
)
