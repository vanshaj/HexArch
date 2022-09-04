package ports

/*
	What this tells us for our adapter to connect to this port it must implement
	1. Addition 2. Subtraction 3. Multiplication 4. Division
*/

type ArithmeticPort interface {
	Addition(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
}
