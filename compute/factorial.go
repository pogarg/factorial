package compute

// Method to calculate factorial of any number
func Factorial(fact chan<- int64, n int) {

	var factorialValue int64 = 1
	if n == 0 {
		factorialValue = 1
	} else {
		for i := 1; i <= n; i++ {
			factorialValue *= int64(i)
		}
	}
	fact <- factorialValue
}
