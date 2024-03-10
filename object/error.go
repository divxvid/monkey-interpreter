package object

// In a real interpreter this object also contains the stack trace,
//
//line numbers etc for better error reporting
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}
