package calc

// Подсчет суммы чисел в слайсе
func SumNumbers(number []int) int {
	sum := 0
	for _, n := range number {
		sum += n
	}
	return sum
}
