package ports

/*
	This is basically port for out application layer
*/

/*
	Here API stands for Application Programming interface
	It is acting as an intermediary that allows two applications to communicate with each other
*/
type APIPort interface {
	GetAddition(a, b int32) (int32, error)
	GetSubtraction(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivision(a, b int32) (int32, error)
}
