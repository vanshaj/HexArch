package arithmetic

/*
	First adapter that will match our port interface created created in the core/domain layer
*/

// We have created our own adapter type. All the adapters we create will have this Adapter type
type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
