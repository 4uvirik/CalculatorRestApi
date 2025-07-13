package calc

// Подсчет суммы чисел в слайсе
func SumNumbers(number []int) (sum int) {
	for _, n := range number {
		sum += n
	}
	return sum
}
