package shared

type ExecError struct {
	Message string
}

func (dbe *ExecError) Error() string {
	return dbe.Message
}

type DatabaseOpeningError struct {
	Message string
}

func (dbe *DatabaseOpeningError) Error() string {
	return dbe.Message
}

type QeuryError struct {
	Message string
}

func (dbe *QeuryError) Error() string {
	return dbe.Message
}

type ServerError struct {
	Message string
}

func (dbe *ServerError) Error() string {
	return dbe.Message
}
