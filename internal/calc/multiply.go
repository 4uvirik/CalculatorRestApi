package calc

// Подсчет умножения чисел в слайсе
func MultiplyNumbers(number []int) (sum int) {
	sum = 1
	for _, n := range number {
		sum *= n
	}
	return sum
}
