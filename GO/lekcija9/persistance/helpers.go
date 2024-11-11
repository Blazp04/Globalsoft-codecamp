package persistance

type DbConnectionError struct {
	Message string
}

type DbStatementError struct {
	Message string
}

func (e *DbConnectionError) Error() string {
	return e.Message
}

func (e *DbStatementError) Error() string {
	return e.Message
}
