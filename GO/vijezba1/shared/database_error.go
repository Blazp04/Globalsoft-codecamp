package shared

type DbConnectionError struct {
	Message string
}

func (db *DbConnectionError) Error() string {
	return db.Message
}
