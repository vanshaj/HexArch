package ports

/*
	We will store whatever operation we will perform in the history table of the database
*/

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}
