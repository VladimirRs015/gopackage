package fib

// Esta funcion es una ejecucion del funcion
func Fib() int {
	// Internal imp
	return FibClico(5)
}

// Este es una fuccion que devuelve el numero correspodiente a su posicion en la sucesion de fibonacci
func FibClico(n int) int {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}
	return a
}
