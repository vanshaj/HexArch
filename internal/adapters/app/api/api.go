package api

import "github.com/vanshaj/HexArch/internal/ports"

/*
	Now we are in the application layer
	If we think of hex architecture , application layer encompasses the core layer. Hence composition.
	So our application layer has a core layer's port and framework layers port too( Composition not inheritance)
*/

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

// Dependency injection of ports's Arithmetic port and frameworks db port
func NewAdapter(arith ports.ArithmeticPort, db ports.DbPort) *Adapter {
	return &Adapter{arith, db}
}

// Implements all methods of application layer ports

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	err = apia.db.AddToHistory(answer, "addition")
	return answer, err
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)
	err = apia.db.AddToHistory(answer, "subtraction")
	return answer, err
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	err = apia.db.AddToHistory(answer, "division")
	return answer, err
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	err = apia.db.AddToHistory(answer, "multiplication")
	return answer, err
}
